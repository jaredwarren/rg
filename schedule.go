package rg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jaredwarren/rg/db"
	schedule "github.com/jaredwarren/rg/gen/schedule"
	"github.com/robfig/cron"
)

// schedule service example implementation.
// The example methods log the requests and return zero values.
type scheduleSvc struct {
	db     db.ScheduleStore
	cron   *cron.Cron
	logger *log.Logger
}

// CronJob ...
type CronJob struct {
	cron     *cron.Cron
	schedule *schedule.Schedule
}

// Run ...
func (c *CronJob) Run() {
	// TODO: set color!!!
	fmt.Println("SET COLOR:", c.schedule.Color)

	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	sched, _ := specParser.Parse(c.schedule.Cron)
	c.schedule.Next = sched.Next(time.Now()).String()

	// for _, e := range c.cron.Entries() {
	// 	if e.
	// }
	// TODO: fiigure out how to update next
	// c.schedule.Next =
}

// NewJob ...
func NewJob(c *cron.Cron, sch *schedule.Schedule) (*CronJob, error) {
	return &CronJob{
		c,
		sch,
	}, nil
}

// NewSchedule returns the schedule service implementation.
func NewSchedule(d *bolt.DB, logger *log.Logger) (schedule.Service, error) {
	// Setup database
	scheduleDb, err := db.NewScheduleDB(d)
	if err != nil {
		return nil, err
	}

	// Start CRON stuff
	c := cron.New()
	var res []*schedule.Schedule
	if res, err = scheduleDb.FetchAll("SCHEDULE"); err != nil {
		return nil, err
	}
	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	for _, sch := range res {
		j, err := NewJob(c, sch)
		if err != nil {
			logger.Print(err)
		}
		c.AddJob(sch.Cron, j)
		// parse
		sched, err := specParser.Parse(sch.Cron)
		if err != nil {
			logger.Print(err)
		}
		if sched != nil {
			sch.Next = sched.Next(time.Now()).String()
		}
	}
	c.Start()
	// https://godoc.org/github.com/robfig/cron

	return &scheduleSvc{scheduleDb, c, logger}, nil
}

// List all stored bottles
func (s *scheduleSvc) List(ctx context.Context) (res []*schedule.Schedule, err error) {
	s.logger.Print("schedule.list")
	if res, err = s.db.FetchAll("SCHEDULE"); err != nil {
		return nil, err
	}
	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor | cron.Descriptor)
	for _, sch := range res {
		sched, err := specParser.Parse(sch.Cron)
		if err != nil {
			s.logger.Print(err)
		}
		sch.Next = sched.Next(time.Now()).String()
	}
	return
}

// Create new cron schedule
func (s *scheduleSvc) Create(ctx context.Context, p *schedule.SchedulePayload) (res *schedule.Schedule, err error) {
	res = &schedule.Schedule{}
	s.logger.Print("schedule.schedule")
	res = &schedule.Schedule{
		Name:  p.Name,
		Cron:  p.Cron,
		Sound: false,
		Color: p.Color,
	}
	// Add Cron Job
	j, err := NewJob(s.cron, res)
	if err != nil {
		return
	}
	err = s.cron.AddJob(res.Cron, j)
	if err != nil {
		return
	}
	// TODO: validate cron here or see if I can add a reges to schedule_service.go
	id, err := s.db.New("SCHEDULE")
	if err != nil {
		return
	}
	res.ID = id
	err = s.db.Save("SCHEDULE", id, res)
	return
}

// Remove cron schedule
func (s *scheduleSvc) Remove(ctx context.Context, p *schedule.RemovePayload) (err error) {
	s.logger.Print("schedule.remove")
	// TODO: delete from db, and stop cron and restart everything :(
	// no way to remove cron
	err = s.db.Delete("SCHEDULE", p.ID)
	if err != nil {
		return
	}

	s.cron.Stop()

	// Start CRON stuff
	s.cron = cron.New()
	var res []*schedule.Schedule
	if res, err = s.db.FetchAll("SCHEDULE"); err != nil {
		return
	}
	specParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	for _, sch := range res {
		j, err := NewJob(s.cron, sch)
		if err != nil {
			s.logger.Print(err)
		}
		s.cron.AddJob(sch.Cron, j)
		// parse
		sched, _ := specParser.Parse(sch.Cron)
		sch.Next = sched.Next(time.Now()).String()
	}
	s.cron.Start()

	return
}

func (s *scheduleSvc) Update(ctx context.Context, p *schedule.UpdatePayload) (err error) {
	s.logger.Print("schedule.color:", p.Color)
	res := &schedule.Schedule{
		ID:    "current_color",
		Name:  "",
		Cron:  "",
		Sound: false,
		Color: p.Color,
	}
	return s.db.Save("SETTINGS", res.ID, res)
}

func (s *scheduleSvc) Color(ctx context.Context) (res *schedule.Color, err error) {
	r, _ := s.db.Fetch("SETTINGS", "current_color")
	if r != nil && r.Color != "" {
		return &schedule.Color{
			Color: r.Color,
		}, nil
	}
	return &schedule.Color{
		Color: "off",
	}, nil
}

func (s *scheduleSvc) Sound(ctx context.Context, p *schedule.SoundPayload) (err error) {
	s.logger.Print("schedule.sound:", p.Sound)
	res := &schedule.Schedule{
		ID:   "current_sound",
		Name: "",
		Cron: "",
	}
	// TODO: fix this so I don't have to use schedule.Schedule to store settings :(
	if p.Sound {
		res.Color = "on"
	} else {
		res.Color = "off"
	}

	return s.db.Save("SETTINGS", res.ID, res)
}

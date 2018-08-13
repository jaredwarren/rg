package rg

import (
	"context"
	"log"

	"github.com/boltdb/bolt"
	"github.com/jaredwarren/rg/cron"
	"github.com/jaredwarren/rg/db"
	schedule "github.com/jaredwarren/rg/gen/schedule"
	"github.com/jaredwarren/rg/pi"
)

// schedule service example implementation.
// The example methods log the requests and return zero values.
type scheduleSvc struct {
	db     db.ScheduleStore
	cron   *cron.Cron
	logger *log.Logger
	pi     pi.Pi
}

// NewSchedule returns the schedule service implementation.
func NewSchedule(d *bolt.DB, logger *log.Logger) (schedule.Service, error) {
	// Setup database
	scheduleDb, err := db.NewScheduleDB(d)
	if err != nil {
		return nil, err
	}

	// Start PI
	rpi := pi.NewPi()

	// Start CRON stuff
	c := cron.NewCron(rpi)
	var schedules []*schedule.Schedule
	if schedules, err = scheduleDb.FetchAll("SCHEDULE"); err != nil {
		return nil, err
	}
	err = c.RunSchedules(schedules)
	if err != nil {
		return nil, err
	}

	return &scheduleSvc{scheduleDb, c, logger, rpi}, nil
}

// List all stored bottles
func (s *scheduleSvc) List(ctx context.Context) (res []*schedule.Schedule, err error) {
	s.logger.Print("schedule.list")
	if res, err = s.db.FetchAll("SCHEDULE"); err != nil {
		return nil, err
	}

	for _, sch := range res {
		next, _ := cron.GetNext(sch)
		sch.Next = next
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
	err = s.cron.AddJob(res)
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

	// Restart cron
	s.cron.Stop()

	s.cron = cron.NewCron(s.pi)
	var schedules []*schedule.Schedule
	if schedules, err = s.db.FetchAll("SCHEDULE"); err != nil {
		return
	}
	err = s.cron.RunSchedules(schedules)
	if err != nil {
		return
	}

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
	// TODO: update physical led
	for _, led := range s.pi.Leds {
		led.Off()
	}
	if p.Color != "off" {
		led, _ := s.pi.Leds[p.Color]
		if led != nil {
			led.On()
		}
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

// Sound Not currently supported
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

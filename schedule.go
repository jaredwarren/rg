package rg

import (
	"context"
	"fmt"
	"log"
	"time"

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
func NewSchedule(d *bolt.DB, rpi pi.Pi, logger *log.Logger) (schedule.Service, error) {
	// Setup database
	scheduleDb, err := db.NewScheduleDB(d)
	if err != nil {
		return nil, err
	}

	// Start CRON stuff
	c := cron.NewCron(func(sch *schedule.Schedule) {
		rpi.SetColor(sch.Color)
		fmt.Println("SET COLOR:", sch.Color, time.Now().String())
	})
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

// List all stored schedules
func (s *scheduleSvc) List(ctx context.Context) (res []*schedule.Schedule, err error) {
	s.logger.Print("schedule.list")
	if res, err = s.db.FetchAll("SCHEDULE"); err != nil {
		return nil, err
	}

	for _, sch := range res {
		next, _ := cron.GetNext(sch)
		sch.Next = next.String()
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

	// no way to remove cron
	if err = s.db.Delete("SCHEDULE", p.ID); err != nil {
		return
	}

	// Restart cron
	s.cron.Stop()

	s.cron = cron.NewCron(func(sch *schedule.Schedule) {
		s.pi.SetColor(sch.Color)
		fmt.Println("SET COLOR:", sch.Color, time.Now().String())
	})
	var schedules []*schedule.Schedule
	if schedules, err = s.db.FetchAll("SCHEDULE"); err != nil {
		return
	}

	if err = s.cron.RunSchedules(schedules); err != nil {
		return
	}

	return
}

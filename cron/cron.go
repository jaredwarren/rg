package cron

import (
	"time"

	"github.com/jaredwarren/rg/gen/schedule"
	"github.com/jaredwarren/rg/pi"
	"github.com/robfig/cron"
)

// Cron ...
type Cron struct {
	cron *cron.Cron
	pi   pi.Pi
}

// NewCron ...
func NewCron(rpi pi.Pi) *Cron {
	return &Cron{cron.New(), rpi}
}

// func Parse(c string) {
// 	cron.Parse()
// }

// RunSchedules ...
func (c *Cron) RunSchedules(schedules []*schedule.Schedule) error {
	for _, sch := range schedules {
		j, err := NewJob(c.cron, sch, c.pi)
		if err != nil {
			// return err // ignore errors
			continue
		}
		c.cron.AddJob(sch.Cron, j)
		// parse
		cronSch, err := cron.Parse(sch.Cron)
		if err != nil {
			// return err // ignore errors
			continue
		}
		if cronSch != nil {
			sch.Next = cronSch.Next(time.Now()).String()
		}
	}
	c.cron.Start()
	// https://godoc.org/github.com/robfig/cron
	return nil
}

// AddJob ...
func (c *Cron) AddJob(schedule *schedule.Schedule) (err error) {
	j, err := NewJob(c.cron, schedule, c.pi)
	if err != nil {
		return
	}
	err = c.cron.AddJob(schedule.Cron, j)
	if err != nil {
		return
	}
	cronSch, err := cron.Parse(schedule.Cron)
	if err != nil {
		return
	}
	if cronSch != nil {
		schedule.Next = cronSch.Next(time.Now()).String()
	}
	return nil
}

// Stop ...
func (c *Cron) Stop() {
	c.cron.Stop()
}

// GetNext ...
func GetNext(schedule *schedule.Schedule) (string, error) {
	sched, err := cron.Parse(schedule.Cron)
	if err != nil {
		return "", err
	}
	return sched.Next(time.Now()).String(), nil
}

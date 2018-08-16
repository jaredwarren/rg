package cron

import (
	"time"

	"github.com/jaredwarren/rg/gen/schedule"
	"github.com/robfig/cron"
)

// Cron wrapper for https://godoc.org/github.com/robfig/cron
type Cron struct {
	cron   *cron.Cron
	runner func(*schedule.Schedule)
}

// NewCron new cron job runner
func NewCron(runner func(*schedule.Schedule)) *Cron {
	return &Cron{cron.New(), runner}
}

// RunSchedules run jobs from schedule list
func (c *Cron) RunSchedules(schedules []*schedule.Schedule) error {
	for _, sch := range schedules {
		j, err := NewJob(c.cron, sch, c.runner)
		if err != nil {
			// ignore errors
			continue
		}
		c.cron.AddJob(sch.Cron, j)
		// parse
		cronSch, err := cron.Parse(sch.Cron)
		if err != nil {
			// ignore errors
			continue
		}
		if cronSch != nil {
			sch.Next = cronSch.Next(time.Now()).String()
		}
	}
	c.cron.Start()
	return nil
}

// AddJob adds cron job to list
func (c *Cron) AddJob(schedule *schedule.Schedule) (err error) {
	j, err := NewJob(c.cron, schedule, c.runner)
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

// Stop cron jobs
func (c *Cron) Stop() {
	c.cron.Stop()
}

// GetNext returns next scheduled run from now
func GetNext(schedule *schedule.Schedule) (time.Time, error) {
	sched, err := cron.Parse(schedule.Cron)
	if err != nil {
		return time.Now(), err
	}
	return sched.Next(time.Now()), nil
}

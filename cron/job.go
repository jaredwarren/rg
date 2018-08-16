package cron

import (
	"time"

	"github.com/jaredwarren/rg/gen/schedule"
	"github.com/robfig/cron"
)

// Job ...
type Job struct {
	cron     *cron.Cron
	schedule *schedule.Schedule
	runner   func(*schedule.Schedule)
}

// NewJob ...
func NewJob(c *cron.Cron, sch *schedule.Schedule, runner func(*schedule.Schedule)) (*Job, error) {
	return &Job{
		c,
		sch,
		runner,
	}, nil
}

// Run ...
func (j *Job) Run() {
	j.runner(j.schedule)
	// calculate next
	schedule, _ := cron.Parse(j.schedule.Cron)
	j.schedule.Next = schedule.Next(time.Now()).String()
}

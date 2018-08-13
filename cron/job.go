package cron

import (
	"fmt"
	"time"

	"github.com/jaredwarren/rg/gen/schedule"
	"github.com/jaredwarren/rg/pi"
	"github.com/robfig/cron"
)

// Job ...
type Job struct {
	cron     *cron.Cron
	schedule *schedule.Schedule
	pi       pi.Pi
}

// NewJob ...
func NewJob(c *cron.Cron, sch *schedule.Schedule, pi pi.Pi) (*Job, error) {
	return &Job{
		c,
		sch,
		pi,
	}, nil
}

// Run ...
func (j *Job) Run() {
	j.pi.SetColor(j.schedule.Color)
	fmt.Println("SET COLOR:", j.schedule.Color, time.Now().String())
	// calculate next
	schedule, _ := cron.Parse(j.schedule.Cron)
	j.schedule.Next = schedule.Next(time.Now()).String()
}

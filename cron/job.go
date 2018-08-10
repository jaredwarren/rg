package cron

import (
	"fmt"
	"time"

	"github.com/jaredwarren/rg/gen/schedule"
	"github.com/jaredwarren/rg/pi"
	"github.com/robfig/cron"
)

// CronJob ...
type CronJob struct {
	cron     *cron.Cron
	schedule *schedule.Schedule
	pi       pi.Pi
}

// NewJob ...
func NewJob(c *cron.Cron, sch *schedule.Schedule, pi pi.Pi) (*CronJob, error) {
	return &CronJob{
		c,
		sch,
		pi,
	}, nil
}

// Run ...
func (j *CronJob) Run() {
	for _, led := range j.pi.Leds {
		led.Off()
	}
	if j.schedule.Color != "off" {
		led, _ := j.pi.Leds[j.schedule.Color]
		if led != nil {
			led.On()
		}
	}

	fmt.Println("SET COLOR:", j.schedule.Color, time.Now().String())
	sched, _ := cron.Parse(j.schedule.Cron)
	j.schedule.Next = sched.Next(time.Now()).String()
}

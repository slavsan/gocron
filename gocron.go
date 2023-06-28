package gocron

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (c *realClock) Now() time.Time {
	return time.Now()
}

type Cron struct {
	clock Clock
}

type Job struct{}

func NewCron() *Cron {
	return NewCronWithClock(&realClock{})
}

func NewCronWithClock(clock Clock) *Cron {
	return &Cron{clock: clock}
}

func (c *Cron) NewJob(title string) error {
	return nil
}

func parse(input string, clock Clock) (time.Duration, time.Duration) {
	now := clock.Now()
	secs := 60 - now.Second() - 1
	millisec := (1000 - (now.UnixMilli() % 1000)) * 1000 * 1000
	startAfter := (time.Duration(secs) * time.Second) + time.Duration(millisec)
	interval := 1 * time.Minute
	return startAfter, interval
}

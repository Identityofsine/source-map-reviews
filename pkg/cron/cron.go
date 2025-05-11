package cron

import (
	"fmt"
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/cron/util"
)

type Executable interface {
	GetName() string
	CronTime() []CronField
	Execute() error
}

type CronJob struct {
	name string
	// Name of the job
	cronTime []CronField
}

func (c *CronJob) CronTime() []CronField {
	return c.cronTime
}

func (c *CronJob) Execute() error {
	// Implement the logic to execute the job
	// For example, you can use a command line tool or a script
	// to perform the task you want to run periodically
	return nil
}

func (c *CronJob) GetName() string {
	return c.name
}

type cron struct {
	crons []Executable
	alive bool
}

var (
	cronInstance = &cron{}
)

func GetCron() *cron {
	return cronInstance
}

func (c *cron) AddCron(cron Executable) {
	c.crons = append(c.crons, cron)
	if c.alive {
		c.startLoop(cron)
	}
}

func (c *cron) AddCrons(crons []Executable) {
	for _, cron := range crons {
		c.AddCron(cron)
	}
}

func (c *cron) GetCrons() []Executable {
	return c.crons
}

func (c *cron) createLoop(job Executable) func() {
	// Implement the logic to loop through the crons and execute them
	// when their cron time happens to be aligned with the current time.
	// A thread will be created to run a timer for each cron job; go routines
	// have little memory overhead, so we can create a lot of them.

	fields := job.CronTime()
	// Durations
	minutes := CronDuration(fields[0]) * time.Minute
	hours := CronDuration(fields[1]) * time.Hour
	days := CronDuration(fields[2]) * time.Hour * 24
	months := CronDuration(fields[3]) * time.Hour * 24 * 30
	totalDuration := minutes + hours + days + months

	return func() {
		// Implement the logic to execute the cron job
		// For example, you can use a command line tool or a script
		// to perform the task you want to run periodically
		for now := range time.Tick(totalDuration) {
			fmt.Printf("Executing cron job, %s, at %s\n", job.GetName(), now.Format(time.RFC3339))
			job.Execute()
		}
	}
}

func (c *cron) startLoop(cron Executable) {
	go c.createLoop(cron)()
}

func (c *cron) Start() {
	for _, cron := range c.crons {
		c.startLoop(cron)
	}
	c.alive = true
	// let this be threaded and not blocking

}

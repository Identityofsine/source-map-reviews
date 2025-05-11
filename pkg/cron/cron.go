package cron

import (
	"fmt"
	"time"
)

type Executable interface {
	GetName() string
	CronTime() time.Duration
	Execute() error
}

type CronJob struct {
	Name       string
	TickerTime time.Duration
}

func (c *CronJob) CronTime() time.Duration {
	return c.TickerTime
}

func (c *CronJob) Execute() error {
	// Implement the logic to execute the job
	// For example, you can use a command line tool or a script
	// to perform the task you want to run periodically
	return nil
}

func (c *CronJob) GetName() string {
	return c.Name
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

func AddCron(cron Executable) {
	cronInstance.crons = append(cronInstance.crons, cron)
	if cronInstance.alive {
		cronInstance.startLoop(cron)
	}
}

func AddCrons(crons []Executable) {
	for _, cron := range crons {
		AddCron(cron)
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

	return func() {
		// Implement the logic to execute the cron job
		// For example, you can use a command line tool or a script
		// to perform the task you want to run periodically
		for now := range time.Tick(job.CronTime() * time.Second) {
			fmt.Printf("Executing cron job, %s, at %s\n", job.GetName(), now.Format(time.RFC3339))
			job.Execute()
		}
	}
}

func (c *cron) startLoop(cron Executable) {
	fmt.Printf("Starting cron job, %s, at %s\n", cron.GetName(), time.Now().Format(time.RFC3339))
	go c.createLoop(cron)()
}

func (c *cron) Start() {
	for _, cron := range c.crons {
		c.startLoop(cron)
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cron"
)

type TestCron struct {
}

func (c *TestCron) GetName() string {
	return "TestCron"
}

func (c *TestCron) CronTime() time.Duration {
	return time.Duration(30)
}

func (c *TestCron) Execute() error {
	fmt.Printf("Executing TestCron at %s\n", time.Now().Format(time.RFC3339))
	return nil
}

func main() {
	fmt.Printf("Hello, World!\n")
	cronInstance := cron.GetCron()

	cronTest := &TestCron{}
	cron.AddCron(cronTest)

	cronInstance.Start()
	select {} // Block forever

}

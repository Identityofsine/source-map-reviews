package main

import (
	"fmt"
	"log"
	"os"
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/cmd/router"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cron"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Starting application...")
	cronInstance := cron.GetCron()

	cronTest := &TestCron{}
	cron.AddCron(cronTest)
	cronInstance.Start()

	router := SetupRouter()
	router.Run(":" + os.Getenv("PORT"))

	select {} // Block forever

}

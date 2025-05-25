package main

import (
	"log"
	"os"

	. "github.com/identityofsine/fofx-go-gin-api-template/cmd/router"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/buildinfo/service"
	cronjobs "github.com/identityofsine/fofx-go-gin-api-template/pkg/cron/jobs"
	cron "github.com/identityofsine/fofx-go-gin-api-template/pkg/cron/services"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	_, derr := db.Connect()
	if derr != nil {
		log.Fatalf("Error connecting to database: %v", derr)
	}
	// Run migrations
	derr = db.Migrate()
	if derr != nil {
		log.Fatalf("Error running migrations: %v", derr.Error())
	}

	if _, err := GetBuildInfo(); err != nil {
		storedlogs.LogFatal("Error getting build info: %v", err)
	}

	storedlogs.LogInfo("Starting application")

	cron.AddJob(cronjobs.GetAuthTokenDeleteJob())

	router := SetupRouter()
	router.Run(":" + os.Getenv("PORT"))

}

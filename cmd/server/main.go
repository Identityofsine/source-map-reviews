package main

import (
	"log"
	"os"

	. "github.com/identityofsine/fofx-go-gin-api-template/cmd/router"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cron"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
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

	log.Println("Starting application...")

	cronInstance := cron.GetCron()
	cronInstance.Start()

	router := SetupRouter()
	router.Run(":" + os.Getenv("PORT"))

}

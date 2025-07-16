package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/identityofsine/fofx-go-gin-api-template/cmd/router"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps/model/mapmodel"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model/mapdb"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/buildinfo/service"
	buildInfoService "github.com/identityofsine/fofx-go-gin-api-template/pkg/buildinfo/service"
	cronjobs "github.com/identityofsine/fofx-go-gin-api-template/pkg/cron/jobs"
	cron "github.com/identityofsine/fofx-go-gin-api-template/pkg/cron/services"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
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

	// Get Build Info
	buildInfo, err := buildInfoService.GetBuildInfo()
	if err != nil {
		log.Fatalf("Error getting build info: %v", err)
	}

	if _, err := GetBuildInfo(); err != nil {
		storedlogs.LogFatal("Error getting build info: %v", err)
	}

	storedlogs.LogInfo(fmt.Sprintf("Starting application under build version %s:<%s> built on: %s\n", buildInfo.Version, buildInfo.CommitHash, buildInfo.BuildDate))

	cron.AddJob(cronjobs.GetAuthTokenDeleteJob())

	mapThing := mapdb.MapDb{
		MapName:   "default_map",
		MapPath:   "/maps/default_map",
		CreatedAt: "2023-10-01T00:00:00Z",
		UpdatedAt: "2023-10-01T00:00:00Z",
	}

	mapp := dbmapper.MapDbFields[mapdb.MapDb, mapmodel.Map](mapThing)
	if mapp == nil {
		fmt.Println("Mapping failed")
	} else {
		fmt.Printf("Mapped Map: %+v\n", *mapp)
	}

	router := SetupRouter()
	router.Run(":" + os.Getenv("PORT"))

}

package services

import (
	"errors"
	"fmt"

	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/cron/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	cron "github.com/robfig/cron/v3"
)

// local/private variable
var (
	cronService = cron.New()
	cronMap     = make(map[string]Executable)
)

func AddJob(job Executable) {
	if job == nil {
		storedlogs.LogFatal("AddJob: job cannot be nil", errors.New("job cannot be nil"))
	}
	if _, exists := cronMap[job.GetName()]; exists {
		storedlogs.LogError("AddJob: job already exists", errors.New("job already exists"))
		return
	}
	storedlogs.LogInfo(fmt.Sprintf("Adding job: %s with cron time: %s", job.GetName(), job.CronTime()))
	cronMap[job.GetName()] = job
	cronService.AddJob(job.CronTime(), job)
	cronService.Start()
}

package jobs

import (
	"fmt"
	"time"

	dbModels "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

type authTokenDeleteJob struct {
}

func (j *authTokenDeleteJob) GetName() string {
	return "AuthTokenDeleteJob"
}

func (j *authTokenDeleteJob) CronTime() string {
	return "* * * * *"
}

func (j *authTokenDeleteJob) Run() {
	// This is to delete expired auth tokens
	dbModels.DeleteTokensWhen("expires_at < timezone('utc', NOW())")
	storedlogs.LogInfo(fmt.Sprintf("Running %s at %s", j.GetName(), time.Now().Format(time.RFC3339)))
}

var (
	jobInstance = &authTokenDeleteJob{}
)

func GetAuthTokenDeleteJob() *authTokenDeleteJob {
	if jobInstance == nil {
		jobInstance = &authTokenDeleteJob{}
	}
	return jobInstance
}

package storedlogs

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/buildinfo/model"
)

//this acts as a mapper

func Map(object BuildInfoDb) BuildInfo {
	return BuildInfo{
		Version:     object.Version,
		CommitHash:  object.CommitHash,
		BuildDate:   object.BuildDate.Format("2006-01-02 15:04:05"),
		Environment: object.Environment,
		CreatedAt:   object.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func MapAll(objects []BuildInfoDb) []BuildInfo {
	logs := make([]BuildInfo, len(objects))
	for i, object := range objects {
		logs[i] = Map(object)
	}
	return logs
}

func ReverseMap(object BuildInfo) BuildInfoDb {
	buildDate, err := time.Parse("2006-01-02 15:04:05", object.BuildDate)
	if err != nil {
		buildDate = time.Time{}
	}
	createdAt, err := time.Parse("2006-01-02 15:04:05", object.CreatedAt)
	if err != nil {
		createdAt = time.Time{}
	}
	return BuildInfoDb{
		Version:     object.Version,
		CommitHash:  object.CommitHash,
		BuildDate:   buildDate,
		Environment: object.Environment,
		CreatedAt:   createdAt,
	}
}

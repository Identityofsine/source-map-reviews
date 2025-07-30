package service

import (
	"time"

	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/health"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/buildinfo/model"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
)

var latestBuildInfo *BuildInfo

func GetBuildInfo() (*BuildInfo, error) {
	if latestBuildInfo != nil {
		return latestBuildInfo, nil
	}
	// Health
	health := GetHealth()
	exists, err := repository.DoesVersionExist(health.Version, health.Commit)
	if err != nil {
		// Handle error
		return nil, err
	}
	if exists {
		// Return the latest build info
		buildInfo, derr := repository.GetBuildInfoByVersionAndCommitHash(health.Version, health.Commit)
		if derr != nil {
			// Handle error
			return nil, derr
		}
		buildInfoObject := dbmapper.MapDbFields[repository.BuildInfoDB, BuildInfo](*buildInfo)
		latestBuildInfo = buildInfoObject
	} else {
		// Create a new build info
		buildInfo := BuildInfo{
			Version:     health.Version,
			CommitHash:  health.Commit,
			BuildDate:   health.BuildDate,
			Environment: health.Environment,
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		}

		buildInfoObject := dbmapper.MapDbFields[BuildInfo, repository.BuildInfoDB](buildInfo)
		err = repository.InsertBuildInfo(*buildInfoObject)
		if err != nil {
			return nil, err
		}
		latestBuildInfo = &buildInfo
	}

	return latestBuildInfo, nil
}

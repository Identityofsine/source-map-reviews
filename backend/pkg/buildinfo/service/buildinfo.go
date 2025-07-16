package service

import (
	"time"

	buildInfoProvider "github.com/identityofsine/fofx-go-gin-api-template/api/dto/buildinfo"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/components/health/service"
	dto "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/buildinfo/model"
)

var latestBuildInfo *BuildInfo

func GetBuildInfo() (*BuildInfo, error) {
	if latestBuildInfo != nil {
		return latestBuildInfo, nil
	}
	// Health
	health := GetHealth()
	exists, err := dto.DoesVersionExist(health.Version, health.Commit)
	if err != nil {
		// Handle error
		return nil, err
	}
	if exists {
		// Return the latest build info
		buildInfo, derr := dto.GetBuildInfoByVersionAndCommitHash(health.Version, health.Commit)
		if derr != nil {
			// Handle error
			return nil, derr
		}
		buildInfoObject := buildInfoProvider.Map(buildInfo)
		latestBuildInfo = &buildInfoObject
	} else {
		// Create a new build info
		buildInfo := BuildInfo{
			Version:     health.Version,
			CommitHash:  health.Commit,
			BuildDate:   health.BuildDate,
			Environment: health.Environment,
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		}

		buildInfoObject := buildInfoProvider.ReverseMap(buildInfo)
		err = dto.InsertBuildInfo(buildInfoObject)
		if err != nil {
			return nil, err
		}
		latestBuildInfo = &buildInfo
	}

	return latestBuildInfo, nil
}

package service

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/model"
)

func GetStoredLogs() []Log {
	// This function should return a slice of Log objects
	logsDbs, err := GetLogs()
	if err != nil {
		storedlogs.LogError("Error getting logs: %v", err)
		return nil
	}
	logs := dbmapper.MapAllDbFields[LogDB, Log](logsDbs)

	return *logs
}

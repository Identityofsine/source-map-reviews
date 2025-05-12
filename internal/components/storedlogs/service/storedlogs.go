package service

import (
	logDto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/storedlogs"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/model"
)

func GetStoredLogs() []Log {
	// This function should return a slice of Log objects
	logsDbs, err := GetLogs()
	if err != nil {
		return nil
	}
	logs := logDto.MapAll(logsDbs)

	return logs
}

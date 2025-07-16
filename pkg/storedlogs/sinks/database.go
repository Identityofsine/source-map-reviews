package sinks

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/api/dto/storedlogs"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/model"
)

type databaseSink struct {
}

func (d *databaseSink) StoreLog(log Log) error {
	logDb := ReverseMap(log)
	err := SaveLogs(logDb)
	return err
}

var DatabaseSink = &databaseSink{}

package sinks

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/db/dbmapper"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/model"
)

type databaseSink struct {
}

func (d *databaseSink) StoreLog(log Log) error {
	logDb := dbmapper.MapDbFields[Log, LogDB](log)
	err := SaveLogs(*logDb)
	return err
}

var DatabaseSink = &databaseSink{}

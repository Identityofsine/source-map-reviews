package storedlogs

import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/model"
)

//this acts as a mapper

func Map(object LogDB) Log {
	return Log{
		ID:        object.Id,
		Severity:  object.Severity,
		Message:   object.Message,
		Timestamp: object.CreatedAt,
	}
}

func MapAll(objects []LogDB) []Log {
	logs := make([]Log, len(objects))
	for i, object := range objects {
		logs[i] = Map(object)
	}
	return logs
}

func ReverseMap(object Log) LogDB {
	return LogDB{
		Id:        object.ID,
		Severity:  object.Severity,
		Message:   object.Message,
		CreatedAt: object.Timestamp,
	}
}

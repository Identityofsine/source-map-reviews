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

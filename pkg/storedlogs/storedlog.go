package storedlogs

import (
	"fmt"
	"time"

	l "github.com/charmbracelet/log"
	cfg "github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs/sinks"
)

//this web application is designed to store logs in a database
//and have them be retrievable via a web interface (for admins)
//and via a REST API (for users)

var (
	sinks = []LogSink{DatabaseSink}
)

func sendToSinks(log Log) {
	log.Version = cfg.GetServerDetails().Version
	log.Commit = cfg.GetServerDetails().Commit
	for _, sink := range sinks {
		err := sink.StoreLog(log)
		if err != nil {
			// Handle error
			continue
		}
	}
}

func LogInfo(log string) {
	logObject := Log{
		Severity:  "INFO",
		Message:   log,
		Timestamp: time.Now(),
	}
	sendToSinks(logObject)
	l.Infof("[%s] INFO: %s\n", logObject.Timestamp.Format(time.RFC3339), log)
}

func LogDebug(log string) {
	logObject := Log{
		Severity:  "DEBUG",
		Message:   log,
		Timestamp: time.Now(),
	}
	sendToSinks(logObject)
	l.Debugf("[%s] DEBUG: \n"+log, logObject.Timestamp.Format(time.RFC3339))
}

func LogWarn(log string) {
	logObject := Log{
		Severity:  "WARN",
		Message:   log,
		Timestamp: time.Now(),
	}
	sendToSinks(logObject)
	l.Warnf("[%s] WARN: \n"+log, logObject.Timestamp.Format(time.RFC3339))
}

func LogError(log string, err error) {
	logObject := Log{
		Severity:  "ERROR",
		Message:   log,
		Timestamp: time.Now(),
	}
	sendToSinks(logObject)
	//get stack trace
	l.Errorf("[%s] ERROR: \n"+log, logObject.Timestamp.Format(time.RFC3339), err)
}

func LogFatal(log string, err error) {
	logObject := Log{
		Severity:  "FATAL",
		Message:   fmt.Sprintf(log, err),
		Timestamp: time.Now(),
	}
	sendToSinks(logObject)
	l.Fatalf("[%s] FATAL: \n"+log, logObject.Timestamp.Format(time.RFC3339), err)
	panic(fmt.Sprintf("[%s] FATAL: "+log, logObject.Timestamp.Format(time.RFC3339)))

}

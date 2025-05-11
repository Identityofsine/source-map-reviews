package util

import "time"

func CronDuration(cronField CronField) time.Duration {
	if cronField.RunEvery == -1 {
		return time.Duration(1)
	} else {
		return time.Duration(cronField.RunEvery)
	}
}

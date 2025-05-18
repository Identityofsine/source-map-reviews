package model

import "time"

type Log struct {
	ID        int       `json:"id"`
	Severity  string    `json:"severity"` //lk
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Commit    string    `json:"commit"`
}

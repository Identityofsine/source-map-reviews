package model

import "time"

type Log struct {
	ID        int       `json:"id" db:"id"`
	Severity  string    `json:"severity" db:"severity"` //lk
	Message   string    `json:"message" db:"message"`
	Timestamp time.Time `json:"timestamp" db:"created_at"`
	Version   string    `json:"version" db:"version"`    //lk
	Commit    string    `json:"commit" db:"commit_hash"` //lk
}

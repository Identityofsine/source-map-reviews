package db

import "fmt"

type dberror struct {
	Source  string `json:"source"`
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
	Major   bool   `json:"major"`
}

type DatabaseError = *dberror

func (e dberror) Error() string {
	return fmt.Sprintf("[%s] %s: %s (%d)", e.Source, e.Message, e.Err, e.Code)
}

func NewDatabaseError(source, message, err string, code int) DatabaseError {
	return &dberror{
		Source:  source,
		Message: message,
		Err:     err,
		Code:    code,
		Major:   false,
	}
}

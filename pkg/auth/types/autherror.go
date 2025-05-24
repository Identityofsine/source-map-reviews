package types

import "fmt"

type autherror struct {
	Source  string `json:"source"`
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

type AuthError = *autherror

func (e *autherror) Error() string {
	return fmt.Sprintf("AUTHERROR: %s: %s (code: %d, error: %s)", e.Source, e.Message, e.Code, e.Err)
}

func NewAuthError(source, message, err string, code int) AuthError {
	return &autherror{
		Source:  source,
		Message: message,
		Err:     err,
		Code:    code,
	}
}

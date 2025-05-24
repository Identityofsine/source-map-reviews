package types

type autherror struct {
	Source  string `json:"source"`
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

type AuthError = *autherror

func NewAuthError(source, message, err string, code int) AuthError {
	return &autherror{
		Source:  source,
		Message: message,
		Err:     err,
		Code:    code,
	}
}

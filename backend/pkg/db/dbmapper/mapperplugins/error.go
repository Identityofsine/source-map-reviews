package mapperplugins

type mappererror struct {
	Source  string `json:"source"`
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

type MapperError = *mappererror

func (e mappererror) Error() string {
	return e.Err
}

func NewMapperError(source, message, err string, code int) MapperError {
	return &mappererror{
		Source:  source,
		Message: message,
		Err:     err,
		Code:    code,
	}
}

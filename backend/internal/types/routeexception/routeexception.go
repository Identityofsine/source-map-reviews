package routeexception

import "fmt"

type routeerror struct {
	Exception error
	Message   string `json:"message"`
	Err       string `json:"error"`
	Code      int    `json:"code"`
}

type RouteError = *routeerror

func (e *routeerror) Error() string {
	return fmt.Sprintf("%s: %s (%d), %s", e.Message, e.Err, e.Code, e.Exception.Error())
}

func NewRouteError(exception error, message, err string, code int) RouteError {
	return &routeerror{
		Exception: exception,
		Message:   message,
		Err:       err,
		Code:      code,
	}
}

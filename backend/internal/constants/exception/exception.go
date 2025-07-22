package exception

import "github.com/identityofsine/fofx-go-gin-api-template/internal/types/routeexception"

var (
	ResourceNotFound = routeexception.NewRouteError(
		nil,
		"Resource not found",
		"resource-not-found",
		404,
	)
)

const (
	CODE_BAD_REQUEST           = 400
	CODE_UNAUTHORIZED          = 401
	CODE_FORBIDDEN             = 403
	CODE_RESOURCE_NOT_FOUND    = 404
	CODE_TOKEN_EXPIRED         = 419
	CODE_TOO_MANY_REQUESTS     = 429
	CODE_NOT_IMPLEMENTED       = 501
	CODE_INTERNAL_SERVER_ERROR = 500
)

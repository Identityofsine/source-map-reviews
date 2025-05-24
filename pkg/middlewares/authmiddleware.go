package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

type authmiddleware struct {
	Middleware gin.HandlerFunc
}

// CreateCors is a constructor function that creates a new CORS middleware; taking in a list of valid origins
func UseAuthenticationEnforcementMiddleware() *authmiddleware {
	return &authmiddleware{
		Middleware: validateToken,
	}
}

func validateToken(c *gin.Context) {
	storedlogs.LogInfo("Validating token")
	c.Next()
}

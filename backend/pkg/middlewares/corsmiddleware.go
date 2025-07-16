package middlewares

import (
	corslib "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	corsdto "github.com/identityofsine/fofx-go-gin-api-template/api/dto/cors"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/config"
)

// Cors is a middleware that sets the
// Access-Control-Allow-Origin header to allow cross-origin requests
// We are able to use this middleware because we are using gin-gonic
// which is a web framework for Go

// CORS is important for web applications that need to make
// requests to a different domain (which is common in microservices and distributed systems)
// It's also important for security, as it helps to prevent cross-site request forgery (CSRF) attacks

// CORS is a security feature implemented by web browsers to prevent malicious websites from making requests to a different domain

type cors struct {
	Middleware gin.HandlerFunc
}

// CreateCors is a constructor function that creates a new CORS middleware; taking in a list of valid origins
func UseCors() *cors {
	config := corsdto.GetCorsConfigFromYaml(*config.GetCorsConfig())
	return &cors{
		Middleware: corslib.New(config),
	}
}

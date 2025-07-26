package middlewares

import (
	"time"

	corslib "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	config := config.GetCorsConfig()
	configObject := getCorsConfigFromYaml(*config)
	return &cors{
		Middleware: corslib.New(configObject),
	}
}

func getCorsConfigFromYaml(config config.CorsConfig) corslib.Config {
	return corslib.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		ExposeHeaders:    config.ExposeHeaders,
		AllowCredentials: config.AllowCredentials,
		MaxAge:           time.Duration(config.MaxAge),
	}
}

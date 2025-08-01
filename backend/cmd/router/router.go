package router

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/auth"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/health"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/images"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/maps"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/register"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/storedlogs"

	"github.com/identityofsine/fofx-go-gin-api-template/pkg/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set up the router
	setupRoutes(r)

	return r
}

func setupRoutes(engine *gin.Engine) {
	// Set up the routes for the application
	api := engine.Group("/api/v1")

	images.ImageRoute.UseRouter(api)

	// Top Level Middleware
	api.Use(middlewares.UseCors().Middleware)

	//inject your routes here:
	//login
	auth.AuthRoute.UseRouter(api)
	register.RegisterRoute.UseRouter(api)
	maps.MapRoute.UseRouter(api)

	api.Use(middlewares.UseAuthenticationEnforcementMiddleware().Middleware)

	health.HealthRoute.UseRouter(api)
	storedlogs.LogsRoute.UseRouter(api)
}

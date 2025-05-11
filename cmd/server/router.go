package main

import (
	"github.com/gin-gonic/gin"
	healthController "github.com/identityofsine/fofx-go-gin-api-template/internal/components/health/controller"
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
	(&healthController.Route{}).UseRouter(api)
}

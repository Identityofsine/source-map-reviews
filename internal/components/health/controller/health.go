package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/health/service"
)

// GET /api/v1/health
func GetServerHealth(c *gin.Context) {
	// Get the environment variable from the request context
	serverHealth := service.GetHealth()
	c.JSON(200, serverHealth)
	return
}

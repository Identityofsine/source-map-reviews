package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/internal/components/storedlogs/service"
)

// GET /api/v1/health
func GetLogs(c *gin.Context) {
	// Get the environment variable from the request context
	logs := service.GetStoredLogs()
	c.JSON(200, logs)
	return
}

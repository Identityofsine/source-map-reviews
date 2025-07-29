package storedlogs

import (
	"github.com/gin-gonic/gin"
)

// GET /api/v1/health
func GetLogs(c *gin.Context) {
	// Get the environment variable from the request context
	logs := GetStoredLogs()
	c.JSON(200, logs)
	return
}

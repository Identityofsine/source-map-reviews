package health

import (
	"github.com/gin-gonic/gin"
)

// GET /api/v1/health
func getServerHealth(c *gin.Context) {
	// Get the environment variable from the request context
	serverHealth := getHealth()
	c.JSON(200, serverHealth)
	return
}

package register

import (
	"strings"

	"github.com/gin-gonic/gin"
	registerPkg "github.com/identityofsine/fofx-go-gin-api-template/pkg/register"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/register/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func genericRegisterHandler(provider registerPkg.Registerable, c *gin.Context) {
	storedlogs.LogInfo("POST: /register/" + strings.ToLower(provider.Name()))

	// Get the request body
	var requestBody RegisterArgs
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the provider's Authenticate method
	if derr := provider.Register(requestBody); derr != nil {
		c.JSON(401, gin.H{"message": "There was an error", "error": derr.Error()})
		return
	}

	// If authentication is successful, return a success response
	c.JSON(200, gin.H{"message": "Registeration successful"})
}

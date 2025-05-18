package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func GenericAuthHandler(provider auth.Authenticator, c *gin.Context) {
	storedlogs.LogInfo("POST: /auth/login/" + strings.ToLower(provider.Name()))

	// Get the request body
	var requestBody AuthenticatorArgs
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the provider's Authenticate method
	if !provider.Authenticate(requestBody) {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// If authentication is successful, return a success response
	c.JSON(200, gin.H{"message": "Authenticated successfully"})
}

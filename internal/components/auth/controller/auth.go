package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	AuthConstants "github.com/identityofsine/fofx-go-gin-api-template/internal/constants/auth"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
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
	if token, err := provider.Authenticate(requestBody); err != nil || token == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	} else {
		// If authentication is successful, return a success response
		c.JSON(200, token)
	}
}

// GET: /auth/refresh
func RefershTokenHandler(c *gin.Context) {
	storedlogs.LogInfo("GET: /auth/refresh")
	cookies := cookies.NewCookies(c)

	userId, err := cookies.GetInt("user_id")
	if err != nil || userId == 0 {
		storedlogs.LogError("RefreshTokenHandler: No user ID found in cookies", err)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	accessToken := c.GetHeader(AuthConstants.AUTHORIZATION_HEADER)

	refreshToken, err := cookies.Get("refresh_token")
	if err != nil || refreshToken == "" {
		storedlogs.LogError("RefreshTokenHandler: No refresh token found in cookies", err)
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	token := Token{
		UserId:       int64(userId),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(200, token)

}

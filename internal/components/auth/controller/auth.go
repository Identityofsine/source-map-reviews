package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	AuthConstants "github.com/identityofsine/fofx-go-gin-api-template/internal/constants/auth"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/model"
	AuthService "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/service"
	. "github.com/identityofsine/fofx-go-gin-api-template/pkg/auth/types"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/cookies"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/storedlogs"
)

func GenericAuthHandler(provider auth.Authenticator, c *gin.Context) {
	storedlogs.LogInfo("POST: /auth/login/" + strings.ToLower(provider.Name()))
	cookies := cookies.NewCookies(c)

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
		AuthService.StoreTokenIntoCookies(*token, cookies)
		c.JSON(200, token)
	}
}

// GET: /auth/refresh
func RefershTokenHandler(c *gin.Context) {
	storedlogs.LogInfo("GET: /auth/refresh")
	cookies := cookies.NewCookies(c)

	userId, err := cookies.GetInt("user_id")
	if err != nil || userId == 0 {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	accessToken := c.GetHeader(AuthConstants.AUTHORIZATION_HEADER)
	if accessToken == "" {
		c.JSON(400, gin.H{"error": "Authorization header is required"})
		return
	}

	refreshToken, err := cookies.Get("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	token := Token{
		UserId:       int64(userId),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if token, err = AuthService.RenewLoginToken(token); err != nil {
		storedlogs.LogError("Error renewing login token", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	AuthService.StoreTokenIntoCookies(token, cookies)

	c.JSON(200, token)

}

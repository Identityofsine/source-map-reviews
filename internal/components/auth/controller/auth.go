package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	UserService "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/service"
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

	user, err := UserService.GetUserByCookies(cookies)
	if err != nil || user == nil {
		c.JSON(401, gin.H{"message": "Unauthorized", "error": "not-authorized"})
		return
	}

	accessToken := c.GetHeader(AuthConstants.AUTHORIZATION_HEADER) // Remove "Bearer " prefix
	accessToken, aerr := AuthService.GetAccessTokenFromHeader(accessToken)
	if accessToken == "" || aerr != nil {
		c.JSON(aerr.Code, gin.H{"message": "Authorization header is required", "error": aerr.Error})
		return
	}

	refreshToken, err := cookies.Get("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(401, gin.H{"error": "unauthorized", "message": "Refresh token is required"})
		return
	}

	token := &Token{
		UserId:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	aerr = nil
	if token, aerr = AuthService.RenewLoginToken(*token, *user); aerr != nil {
		c.JSON(aerr.Code, gin.H{"error": aerr.Err, "message": aerr.Message})
		return
	}

	AuthService.StoreTokenIntoCookies(*token, cookies)

	c.JSON(200, token)

}

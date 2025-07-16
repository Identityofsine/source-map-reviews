package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	UserService "github.com/identityofsine/fofx-go-gin-api-template/internal/components/user/service"
	AuthConstants "github.com/identityofsine/fofx-go-gin-api-template/internal/constants/auth"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/repository/model"
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
	// check if the request is a post or get request
	if c.Request.Method != "POST" && c.Request.Method != "GET" {
		c.JSON(405, gin.H{"error": "Method not allowed", "message": "Only POST and GET methods are allowed"})
		return
	}

	var requestContent map[string]interface{}
	if c.Request.Method == "POST" {
		if err := c.ShouldBindJSON(&requestContent); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
	} else {
		// For GET requests, we can use query parameters
		requestContent = make(map[string]interface{})
		for key, values := range c.Request.URL.Query() {
			if len(values) > 0 {
				requestContent[key] = values[0] // Use the first value for simplicity
			}
		}
	}

	requestBody := NewAuthenticatorArgs()
	requestBody.Keys = requestContent
	requestBody.Context = c

	// Call the provider's Authenticate method
	if token, err := provider.Authenticate(requestBody); err != nil {
		c.JSON(err.Code, err)
		return
	} else if token != nil {
		// If authentication is successful, return a success response
		AuthService.StoreTokenIntoCookies(*token, cookies)
		c.JSON(200, token)
	} else {
		//assume that the Authenticator ended up redirecting
		return
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

func LogoutHandler(c *gin.Context) {
	storedlogs.LogInfo("POST: /auth/logout")
	cookies := cookies.NewCookies(c)

	// Get the user from cookies
	user, err := UserService.GetUserByCookies(cookies)
	if err != nil || user == nil {
		c.JSON(401, gin.H{"message": "Unauthorized", "error": "not-authorized"})
		return
	}

	// Get the token from cookies
	token, aerr := AuthService.GetTokenFromCookies(cookies)
	if aerr != nil {
		c.JSON(aerr.Code, gin.H{"message": aerr.Message, "error": aerr.Err})
		return
	}

	// Delete the token from the database
	derr := DeleteTokenByRefreshToken(token.RefreshToken)
	if derr != nil {
		storedlogs.LogError("LogoutHandler: Failed to delete token from database", err)
		c.JSON(derr.Code, gin.H{"message": "Failed to delete token from database", "error": err.Error()})
		return
	}

	// Delete the token
	if derr := AuthService.DeleteTokenInCookies(cookies); derr != nil {
		c.JSON(500, gin.H{"message": "Failed to delete token", "error": derr.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

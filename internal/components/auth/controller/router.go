package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/auth"
)

type route Routeable

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	authProviders := auth.GetAuthProviders()
	authRoute := router.Group("/auth")
	loginGroup := authRoute.Group("/login")
	{
		for _, provider := range authProviders {
			provider := provider
			providerName := strings.ToLower(provider.Name())
			loginGroup.POST("/"+providerName, func(c *gin.Context) {
				GenericAuthHandler(provider, c)
			})
		}
	}

	//refresh
	authRoute.GET("/refresh", RefershTokenHandler)

	//logout
	authRoute.GET("/logout", LogoutHandler)

	return router
}

var (
	AuthRoute = route{}
)

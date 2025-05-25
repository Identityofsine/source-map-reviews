package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
	"github.com/identityofsine/fofx-go-gin-api-template/pkg/register"
)

type route Routeable

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	registerProviders := register.GetRegisterProviders()
	registerGroup := router.Group("/register")
	{
		for _, provider := range registerProviders {
			provider := provider
			providerName := strings.ToLower(provider.Name())
			registerGroup.POST("/"+providerName, func(c *gin.Context) {
				GenericRegisterHandler(provider, c)
			})
		}
	}

	return router
}

var (
	RegisterRoute = route{}
)

package images

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	registerGroup := router.Group("/images")
	registerGroup.GET("/:path", GetImageRoute)

	return router
}

var (
	ImageRoute = route{}
)

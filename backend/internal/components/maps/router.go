package maps

import (
	"github.com/gin-gonic/gin"
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	registerGroup := router.Group("/maps")
	registerGroup.GET("/", GetMapsRoute)
	registerGroup.GET("/:mapName", GetMapRoute)
	registerGroup.POST("/search", SearchMapsRoute)

	return router
}

var (
	MapRoute = route{}
)

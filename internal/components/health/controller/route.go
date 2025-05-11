package controller

import "github.com/gin-gonic/gin"
import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types"
)

type Route Routeable

//this directory contains the routes responsible for handling the requests
//of the health component of this web application

func (_ *Route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("/health", GetServerHealth)
	return router
}

package controller

import "github.com/gin-gonic/gin"
import (
	. "github.com/identityofsine/fofx-go-gin-api-template/internal/types/router"
)

type route Routeable

//this directory contains the routes responsible for handling the requests
//of the health component of this web application

func (_ *route) UseRouter(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("/logs", GetLogs)
	return router
}

var (
	LogsRoute = route{}
)

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quentinguidee/microservice-core/router"
)

func initializeRouter() *gin.Engine {
	r := router.CreateRouter()

	return r
}

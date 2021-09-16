package main

import (
	. "github.com/971181317/chaos-pic/action"
	"github.com/gin-gonic/gin"
)

func RegisterRoute() *gin.Engine {
	route := gin.Default()
	gin.Logger()

	chaosGroup := route.Group("/chaos-pic")
	{
		chaosGroup.POST("/upload", FileUpload)
	}

	return route
}

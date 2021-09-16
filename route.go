package main

import (
	. "github.com/971181317/chaos-pic/action"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute() *gin.Engine {
	route := gin.New()

	//注册中间件
	route.Use(InitReq)
	route.Use(gin.Recovery())

	route.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/chaos-pic")
	})

	chaosPic := route.Group("/chaos-pic")
	{
		chaosPic.POST("/upload", FileUpload)
		chaosPic.GET("/upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{})
		})
	}

	return route
}

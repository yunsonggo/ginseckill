package router

import (
	"2022/ginseckill/middleware"
	"2022/ginseckill/web/webAuthController"
	"github.com/gin-gonic/gin"
)

func WebAuthGroup(webGroup  *gin.RouterGroup) {
	webAuthGroup := webGroup.Group("/auth")
	webAuthGroup.Use(middleware.WebAuth())
	{
		webAuthGroup.GET("/ping",webAuthController.Ping)
	}
}
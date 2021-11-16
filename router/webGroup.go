package router

import (
	"2022/ginseckill/web/webController"
	"github.com/gin-gonic/gin"
)

func WebGroup(app *gin.Engine) {
	webGroup := app.Group("/api")
	{
		webGroup.GET("/ping",webController.Ping)
		// auth
		WebAuthGroup(webGroup)
	}
}
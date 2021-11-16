package router

import (
	"2022/ginseckill/backend/backendController"
	"github.com/gin-gonic/gin"
)

func BackendGroup(app *gin.Engine) {
	bg := app.Group("/manager")
	{
		bg.GET("/ping",backendController.Ping)
		// 注册用户
		bg.POST("/sign/up",backendController.InsertConsumer)
		// auth
		BackendAuthGroup(bg)
	}
}

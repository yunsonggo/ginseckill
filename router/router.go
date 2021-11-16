package router

import (
	"2022/ginseckill/common"
	"2022/ginseckill/config"
	"2022/ginseckill/middleware"
	"github.com/gin-gonic/gin"
	"github.com/yunsonggo/loggo"
	"net/http"
)

func NewRouter() *gin.Engine {
	wc := config.Conf.Website
	if wc.WebsiteMode == gin.ReleaseMode {
		loggo.Init(loggo.Config{
			Path: wc.WebsiteLog,
			Stdout: wc.Stdout,
		})
		gin.DefaultWriter = loggo.InfoLog
		gin.DefaultErrorWriter = loggo.ErrorLog
	}
	// 初始化引擎
	app := gin.Default()
	// 设置静态文件路径
	app.StaticFS("/api/static",http.Dir("./public"))
	// 启用session
	common.InitSession(app)
	// 启用cors
	app.Use(middleware.CorsMiddleware())
	WebGroup(app)
	BackendGroup(app)
	return app
}

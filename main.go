package main

import (
	"2022/ginseckill/config"
	"2022/ginseckill/dial"
	"2022/ginseckill/router"
	"github.com/yunsonggo/loggo"
)

func main() {
	// 初始化配置文件
	config.ConfigInit()
	// 启用数据库
	err := dial.InitEngines()
	if err != nil {
		loggo.Error(err)
		return
	}
	// 启用路由
	app := router.NewRouter()
	// 启用服务
	if config.Conf.Website.WebsiteMode == "release" {
		loggo.Error(app.Run(config.Conf.Listen))
	} else {
		_ = app.Run(config.Conf.Listen)
	}
}
package middleware

import (
	"2022/ginseckill/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域
func CorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	corsExample := config.Conf.Website.CorsExample
	// 运行Release模式下才会进行跨域保护 保证开发过程中不会被跨域困扰
	if gin.Mode() == gin.ReleaseMode {
		corsConfig.AllowOrigins = []string{corsExample}
	} else {
		corsConfig.AllowAllOrigins = true
	}
	corsConfig.AllowCredentials = true
	return cors.New(corsConfig)
}
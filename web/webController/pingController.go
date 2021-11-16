package webController

import (
	"2022/ginseckill/common"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	common.Success(ctx,"pong")
	return
}

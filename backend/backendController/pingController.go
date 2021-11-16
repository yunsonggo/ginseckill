package backendController

import (
	"2022/ginseckill/common"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	common.Success(ctx,"back pong")
	return
}

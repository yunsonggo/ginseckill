package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0	// 成功代码
	FAILED int = 1 // 失败代码
)

func ResponseMsg(ctx *gin.Context,err error,v interface{}) bool {
	if err != nil {
		Failed(ctx,err)
		return false
	} else {
		Success(ctx,v)
		return true
	}
}

func ResponseErr(ctx *gin.Context,err error) bool {
	if err != nil {
		Failed(ctx,err)
		return false
	}
	return true
}

// 成功返回
func Success(ctx *gin.Context,v interface{}) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":SUCCESS,
		"msg":"ok",
		"data":v,
	})
}
//失败返回
func Failed(ctx *gin.Context,v interface{}) {
	ctx.JSON(http.StatusOK,gin.H{
		"code":FAILED,
		"msg":v,
	})
}
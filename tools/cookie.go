package tools

import (
	"github.com/gin-gonic/gin"
)

//设置全局cookie

func GlobalCookie(ctx *gin.Context,name,value string) {
	ctx.SetCookie(name,value,3600,"/","",false,true)
}

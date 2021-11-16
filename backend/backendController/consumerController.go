package backendController

import (
	"2022/ginseckill/common"
	"2022/ginseckill/models"
	"2022/ginseckill/service"
	"2022/ginseckill/tools"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

var cs = service.NewConsumerServer()

func SelectConsumerByName(ctx *gin.Context) {
	name := ctx.Query("name")
	res, err := cs.SelectByName(name)
	_ = common.ResponseMsg(ctx, err, res)
	return
}

func InsertConsumer(ctx *gin.Context) {
	consumer := new(models.ConsumerModel)
	err := ctx.ShouldBindBodyWith(consumer, binding.JSON)
	if err != nil {
		common.Failed(ctx, err.Error())
		return
	}
	_, err = cs.InsertConsumer(consumer)
	if err != nil {
		common.Failed(ctx, err.Error())
		return
	}
	con,err:= cs.SelectByName(consumer.Name)
	if err != nil {
		common.Failed(ctx, err.Error())
		return
	}
	id := con.Id
	idStr := strconv.FormatInt(id, 10)
	signByte := tools.AesCtrEncrypt([]byte(idStr))
	sign := tools.UrlBase64Encode(signByte)
	tools.GlobalCookie(ctx,"uid",idStr)
	tools.GlobalCookie(ctx,"sign",sign)
	common.Success(ctx, "ok")
	return
}

package backendAuthController

import (
	"2022/ginseckill/common"
	"2022/ginseckill/models"
	"2022/ginseckill/param"
	"2022/ginseckill/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var gs = service.NewGoodsServer()

func Ping(ctx *gin.Context) {
	common.Success(ctx,"b a pong")
	return
}
// 所有商品
func GoodsList(ctx *gin.Context) {
	list,err := gs.SelectAll()
	_ = common.ResponseMsg(ctx,err,list)
	return
}

// 更新商品
func GoodsUpdate(ctx *gin.Context) {
	info := new(models.Goods)
	err := ctx.ShouldBindBodyWith(info,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	err = gs.Update(info)
	_ = common.ResponseMsg(ctx,err,"更新成功!")
	return
}

// 添加商品
func InsertGoods(ctx *gin.Context) {
	info := new(models.Goods)
	err := ctx.ShouldBindBodyWith(info,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	_,err = gs.Insert(info)
	_ = common.ResponseMsg(ctx,err,"添加成功!")
	return
}

// 删除一件商品
func DeleteGoodsOne(ctx *gin.Context) {
	info := new(models.Goods)
	err := ctx.ShouldBindBodyWith(info,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	err = gs.Delete(info.Id)
	_ = common.ResponseMsg(ctx,err,"删除成功!")
	return
}

// 单条件查询
func SelectOneGoods(ctx *gin.Context) {
	con := new(param.Condition)
	err := ctx.ShouldBindBodyWith(con,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	res,err := gs.SelectByField(con.Field,con.Value)
	if !common.ResponseErr(ctx,err) {
		return
	}
	_ = common.ResponseMsg(ctx,nil,res)
	return
}
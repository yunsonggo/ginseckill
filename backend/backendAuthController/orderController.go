package backendAuthController

import (
	"2022/ginseckill/common"
	"2022/ginseckill/models"
	"2022/ginseckill/param"
	"2022/ginseckill/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var os = service.NewOrderServer()

// 所有订单
func OrderList(ctx *gin.Context) {
	list,err := os.SelectAll()
	_ = common.ResponseMsg(ctx,err,list)
	return
}

// 更新订单
func OrderUpdate(ctx *gin.Context) {
	info := new(models.Order)
	err := ctx.ShouldBindBodyWith(info,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	err = os.Update(info)
	_ = common.ResponseMsg(ctx,err,"更新成功!")
	return
}

// 添加订单
func InsertOrder(ctx *gin.Context) {
	info := new(models.Order)
	err := ctx.ShouldBindBodyWith(info,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	_,err = os.Insert(info)
	_ = common.ResponseMsg(ctx,err,"添加成功!")
	return
}

// 删除一件订单
func DeleteOrderOne(ctx *gin.Context) {
	info := new(models.Order)
	err := ctx.ShouldBindBodyWith(info,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	err = os.Delete(info.Id)
	_ = common.ResponseMsg(ctx,err,"删除成功!")
	return
}

// 单条件查询
func SelectOneOrder(ctx *gin.Context) {
	con := new(param.Condition)
	err := ctx.ShouldBindBodyWith(con,binding.JSON)
	if !common.ResponseErr(ctx,err) {
		return
	}
	res,err := os.SelectByField(con.Field,con.Value)
	if !common.ResponseErr(ctx,err) {
		return
	}
	_ = common.ResponseMsg(ctx,nil,res)
	return
}
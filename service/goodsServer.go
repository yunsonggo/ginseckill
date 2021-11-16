package service

import (
	"2022/ginseckill/dial"
	"2022/ginseckill/models"
	"fmt"
	"github.com/yunsonggo/loggo"
	"strconv"
)

type GoodsServer interface {
	// 添加商品
	Insert(goods *models.Goods) (id int64,err error)
	// 删除商品
	Delete(id int64) (err error)
	// 更新
	Update(goods *models.Goods) (err error)
	// 根据field查询一条
	SelectByField(field,value string) (goods *models.Goods,err error)
	// 查询所有
	SelectAll() (goodsList []models.Goods,err error)
}

type goodsServer struct {}

func NewGoodsServer() GoodsServer {
	return &goodsServer{}
}
// 错误日志
func (gs *goodsServer) logErr(err error,msg string,data interface{}) bool {
	if err != nil {
		info := fmt.Sprintf("%+v",data)
		loggo.ErrorFormat("%s: %v\n,err:%v",msg,info,err)
		return false
	}
	return true
}
// 添加商品
func (gs *goodsServer) Insert(goods *models.Goods) (id int64,err error) {
	id,err = dial.D.InsertOne(goods)
	gs.logErr(err,"添加商品失败,商品信息:",&goods)
	return
}
// 删除商品
func (gs *goodsServer) Delete(id int64) (err error) {
	goods := new(models.Goods)
	_,err = dial.D.ID(id).Delete(goods)
	gs.logErr(err,"删除商品失败,商品id:",id)
	return
}
// 更新
func (gs *goodsServer) Update(goods *models.Goods) (err error) {
	_,err = dial.D.ID(goods.Id).Update(goods)
	gs.logErr(err,"更新商品失败,商品信息:",&goods)
	return
}
// 根据field查询一条
func (gs *goodsServer) SelectByField(field,value string) (goods *models.Goods,err error) {
	if field == "id" {
		id,parseErr := strconv.ParseInt(value,10,64)
		if parseErr != nil {
			err = parseErr
			return
		}
		_,err = dial.D.Where("id = ?",id).Get(goods)
	} else {
		_,err = dial.D.Where("? = ?",field,value).Get(goods)
	}
	gs.logErr(err,"查询商品失败,查询field--value:",field+"---"+value)
	return
}
// 查询所有
func (gs *goodsServer) SelectAll() (goodsList []models.Goods,err error) {
	err = dial.D.Find(&goodsList)
	return
}

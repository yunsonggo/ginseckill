package service

import (
	"2022/ginseckill/dial"
	"2022/ginseckill/models"
	"fmt"
	"github.com/yunsonggo/loggo"
	"strconv"
)

type OrderServer interface {
	Insert(order *models.Order) (id int64,err error)
	Delete(id int64) (err error)
	Update(order *models.Order) (err error)
	SelectByField(field,value string) (order *models.Order,err error)
	SelectAll() (orderList []*models.Order,err error)
}

type orderServer struct {}

func NewOrderServer() OrderServer {
	return &orderServer{}
}

// 错误日志
func (os *orderServer) logErr(err error,msg string,data interface{}) bool {
	if err != nil {
		info := fmt.Sprintf("%+v",data)
		loggo.ErrorFormat("%s: %v\n,err:%v",msg,info,err)
		return false
	}
	return true
}

func (os *orderServer) Insert(order *models.Order) (id int64,err error) {
	id,err = dial.D.InsertOne(order)
	os.logErr(err,"添加订单失败,订单:",&order)
	return
}

func (os *orderServer) Delete(id int64) (err error) {
	order := new(models.Order)
	_,err = dial.D.ID(id).Delete(order)
	os.logErr(err,"删除订单失败,订单:",id)
	return
}

func (os *orderServer) Update(order *models.Order) (err error) {
	_,err = dial.D.ID(order.Id).Update(order)
	os.logErr(err,"更新订单失败,订单信息:",&order)
	return
}

func (os *orderServer) SelectByField(field,value string) (order *models.Order,err error) {
	if field == "id" {
		id,parseErr := strconv.ParseInt(value,10,64)
		if parseErr != nil {
			err = parseErr
			return
		}
		_,err = dial.D.Where("id = ?",id).Get(order)
	} else {
		_,err = dial.D.Where("? = ?",field,value).Get(order)
	}
	os.logErr(err,"查询订单失败,查询field--value:",field+"---"+value)
	return
}

func (os *orderServer) SelectAll() (orderList []*models.Order,err error) {
	err = dial.D.Find(orderList)
	return
}
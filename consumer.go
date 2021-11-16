package main

import (
	"2022/ginseckill/models"
	"2022/ginseckill/rabbitmq"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/yunsonggo/loggo"
	"strconv"
)

var (
	consumerDB *xorm.Engine
	conn = "root@tcp(127.0.0.1:3306)/ginseckill?charset=utf8"
	driver = "mysql"
	showdb = true
)

func newMysqlEngine(conn,driver string,show bool) (err error) {
	db,err := xorm.NewEngine(driver,conn)
	if err != nil {
		panic(fmt.Sprintf("连接数据库异常:%v\n",err))
	}
	db.ShowSQL(show)
	consumerDB = db
	return
}

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
	id,err = consumerDB.InsertOne(order)
	os.logErr(err,"添加订单失败,订单:",&order)
	return
}

func (os *orderServer) Delete(id int64) (err error) {
	order := new(models.Order)
	_,err = consumerDB.ID(id).Delete(order)
	os.logErr(err,"删除订单失败,订单:",id)
	return
}

func (os *orderServer) Update(order *models.Order) (err error) {
	_,err = consumerDB.ID(order.Id).Update(order)
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
		_,err = consumerDB.Where("id = ?",id).Get(order)
	} else {
		_,err = consumerDB.Where("? = ?",field,value).Get(order)
	}
	os.logErr(err,"查询订单失败,查询field--value:",field+"---"+value)
	return
}

func (os *orderServer) SelectAll() (orderList []*models.Order,err error) {
	err = consumerDB.Find(orderList)
	return
}

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
	id,err = consumerDB.InsertOne(goods)
	gs.logErr(err,"添加商品失败,商品信息:",&goods)
	return
}
// 删除商品
func (gs *goodsServer) Delete(id int64) (err error) {
	goods := new(models.Goods)
	_,err = consumerDB.ID(id).Delete(goods)
	gs.logErr(err,"删除商品失败,商品id:",id)
	return
}
// 更新
func (gs *goodsServer) Update(goods *models.Goods) (err error) {
	_,err = consumerDB.ID(goods.Id).Update(goods)
	gs.logErr(err,"更新商品失败,商品信息:",&goods)
	return
}
// 根据field查询一条
func (gs *goodsServer) SelectByField(field,value string) (goods *models.Goods,err error) {
	goodInfo := new(models.Goods)
	if field == "id" {
		id,parseErr := strconv.ParseInt(value,10,64)
		if parseErr != nil {
			err = parseErr
			return
		}
		_,err = consumerDB.Where("id = ?",id).Get(goodInfo)
	} else {
		_,err = consumerDB.Where("? = ?",field,value).Get(goodInfo)
	}
	gs.logErr(err,"查询商品失败,查询field--value:",field+"---"+value)
	goods = goodInfo
	return
}
// 查询所有
func (gs *goodsServer) SelectAll() (goodsList []models.Goods,err error) {
	err = consumerDB.Find(&goodsList)
	return
}

func main() {
	_ = newMysqlEngine(conn,driver,showdb)
	var rabbitMQConsumer *rabbitmq.RabbitMQ
	var mqQueueName = "imoocProduct"
	os := NewOrderServer()
	ps := NewGoodsServer()

	rabbitMQConsumer = rabbitmq.NewRabbitMQSimple(mqQueueName)
	rabbitMQConsumer.ConsumerSimple(os,ps)
	fmt.Println("consumer start")
}


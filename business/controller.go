package business

import (
	"2022/ginseckill/models"
	"2022/ginseckill/rabbitmq"
	"encoding/json"
	"fmt"
	"github.com/yunsonggo/loggo"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	getOneIp = "192.168.1.102"
	getOnePort = "8084"
	accessControl = &AccessControl{sourcesArray: make(map[int]time.Time)}
	// 全局变量rabbitMQValidate
	rabbitMQValidata *rabbitmq.RabbitMQ
	mqQueueName = "imoocProduct"
)

// "/checkRight"接口 绑定的控制器 用于其他节点来本机获取访问信息确认
// 用于流量控制 10秒内访问一次
func CheckRight(w http.ResponseWriter,r *http.Request) {
	right := accessControl.GetDistributedRight(r)
	if !right {
		_,_ = w.Write([]byte("false"))
		return
	}
	_,_ = w.Write([]byte("true"))
	return
}

// 正常业务逻辑 生成订单 操作数据库
func Check(w http.ResponseWriter,r *http.Request)  {
	fmt.Printf("业务check")
	// 获取URL参数 `x=1&y=2&y=3;z` to json {"x":["1"], "y":["2", "3"], "z":[""]}
	queryForm,err := url.ParseQuery(r.URL.RawQuery)
	if err != nil || len(queryForm["productID"]) <= 0 {
		_,_ = w.Write([]byte("get pid false"))
		return
	}
	productString := queryForm["productID"][0]
	fmt.Println("商品ID为", productString)
	// 获取用户cookie
	userCookie, err := r.Cookie("uid")
	if err != nil {
		_,_ = w.Write([]byte("cookie uid false"))
		return
	}
	// 分布式权限验证 如果访问信息存储在本机,则验证
	// 这里验证的是访问时间 第一次访问则通过
	// 不是第一次 则10秒后可以再次访问
	// 如果是其他节点 则通过GetCurl()模拟http请求其他节点的 "/checkRight"接口 验证
	right := accessControl.GetDistributedRight(r)
	if right == false {
		_,err = w.Write([]byte("time control false"))
		return
	}
	// 获取数量控制权限， 防止秒杀出现超卖现象
	getOneUrl := "http://"+getOneIp+":"+getOnePort+"/getOne"
	responseValidate,validateBody,err := GetCurl(getOneUrl,r)
	if err != nil {
		_,_ = w.Write([]byte("getOne false"))
		return
	}
	// 判断数量控制接口请求状态
	if responseValidate.StatusCode == 200 {
		if string(validateBody) == "true" {
			// 整合下单
			// 1.获取商品ID
			productID, err := strconv.ParseInt(productString, 10, 64)
			if err != nil {
				ToBackOne(w,r,productID)
				return
			}
			// 2.获取用户ID
			userID, err := strconv.ParseInt(userCookie.Value, 10, 64)
			if err != nil {
				ToBackOne(w,r,productID)
				return
			}
			// 3.创建消息体
			message := models.NewMessage(userID,productID)
			byteMessage, err := json.Marshal(message)
			if err != nil {
				ToBackOne(w,r,productID)
				return
			}
			// 4.RabbitMQ Simple 模式 生产消息
			rabbitMQValidata = rabbitmq.NewRabbitMQSimple(mqQueueName)
			defer rabbitMQValidata.Destory()
			err = rabbitMQValidata.PublishSimple(string(byteMessage))
			if err != nil {
				_,_ = w.Write([]byte("mq publish false"))
				return
			}
			_,err = w.Write([]byte("true"))
			return
		}
	}
	_,err = w.Write([]byte("response status false"))
	return
}

// 退还一个计数
func ToBackOne(w http.ResponseWriter,r *http.Request,pid int64)  {
	backOneUrl := "http://"+getOneIp+":"+getOnePort+"/backOne"
	responseValidate,validateBody,err := GetCurl(backOneUrl,r)
	if err != nil {
		_,_ = w.Write([]byte("to back one false"))
		loggo.ErrorFormat("退还计数getOne: sum -= 1失败,pId:%d",pid)
		return
	}
	if responseValidate.StatusCode == 200 {
		if string(validateBody) == "true" {
			_,err = w.Write([]byte("true"))
			loggo.ErrorFormat("退还计数getOne: sum -= 1成功,pId:%d",pid)
			return
		}
	}
	_,err = w.Write([]byte("to back one false"))
	loggo.ErrorFormat("退还计数getOne: sum -= 1失败,pId:%d",pid)
	return
}
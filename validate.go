package main

import (
	"2022/ginseckill/business"
	"2022/ginseckill/tools"
	"fmt"
	"github.com/yunsonggo/loggo"
	"net/http"
)

// 基于COOKIE的独立中间件权限验证
// 服务器节点哈希环获取节点数据
// 秒杀流量控制
// rabbitMQ队列
// 消费商品生成订单
var (
	localhost string
	port = "8081"
	hostArray = []string{"127.0.0.1","192.168.1.102","0.0.0.0"}
)

func main() {
	var err error
	// 1.获取本机IP 赋值全局变量localhost 用于请求本机服务
	localhost,err = tools.GetIntranceIp()
	if err != nil {
		loggo.ErrorFormat("获取本机IP错误,采用默认IP:%s\n","127.0.0.1")
		localhost = "127.0.0.1"
	}
	fmt.Printf("localhost:%s\n",localhost)
	// 3.负载均衡设置 一致性哈希算法
	// 添加节点 v 是服务器节点IP
	business.AddHashConsistent(localhost,port,hostArray)
	// 4.自定义中间件 验证权限
	filter := business.NewFileter()
	// 注册拦截器 比如要访问的是 "/check",就注册"/check"
	// 访问"/check" 就会先访问这里注册的拦截器 也就是 中间件
	filter.RegisterUri("/check",business.Auth)
	filter.RegisterUri("/checkRight",business.Auth)
	// 流量控制 这里设置10秒内访问一次 用于其他节点来本机获取访问信息确认
	http.HandleFunc("/checkRight",filter.Handle(business.CheckRight))
	// 启用拦截器中间件Auth通过后 执行的正常业务逻辑 类似于 控制器
	http.HandleFunc("/check",filter.Handle(business.Check))
	// 启动服务
	err = http.ListenAndServe(localhost+":"+port,nil)
	fmt.Printf("listenAndServer:%s\n",localhost+":"+port)
	loggo.Error(err)
}

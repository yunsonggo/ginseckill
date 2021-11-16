package rabbitmq

import (
	"2022/ginseckill/config"
	"2022/ginseckill/models"
	"2022/ginseckill/service"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yunsonggo/loggo"
	"log"
	"strconv"
	"sync"
)

// 连接地址
var mqAddr = config.Conf.Rabbitmq.RabbitmqAddr

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	// 队列
	QueueName string
	// 交换机
	Exchange string
	// bind key
	Key string
	// conn
	Mqurl string
	//锁
	sync.Mutex
}

// 基础实例 给各种模式提供实例基础 不同参数 不同模式
func NewRabbitMQ(queueName,exchange,key string) *RabbitMQ {
	if len(mqAddr) == 0 {
		mqAddr = "amqp://root:root@192.168.1.102:5672/imooc"
	}
	fmt.Printf("mqAddr:%s\n",mqAddr)
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: mqAddr}
	var err error
	// 创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	fmt.Printf("conn mq err:%+v\n",err)
	_ = rabbitmq.failOnErr(err, "创建连接错误！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	_ = rabbitmq.failOnErr(err, "获取channel失败！")
	return rabbitmq
}

// 错误处理
func (r *RabbitMQ) failOnErr(err error,message string) error{
	if err != nil {
		loggo.ErrorFormat("%s:%v",message,err)
	}
	return err
}

// 断开 关闭
func (r *RabbitMQ) Destory() (err error) {
	err = r.channel.Close()
	_ = r.failOnErr(err,"close rabbitMQ channel failed")
	err = r.conn.Close()
	_ = r.failOnErr(err,"close rabbitMQ conn failed")
	return
}

// step 一: 简单模式 实例 简单模式下只需要 queueName 其他参数默认 一个生产者 一个消费者
// 工作模式: 与简单模式区别只有: 一个生产者 多个消费者 一个消息只能被一个消费者消费
// 代码不变 (当生产速度大于消费速度时,负载均衡作用)
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}
// step 二: 简单模式生产消息 加锁
func (r *RabbitMQ) PublishSimple(message string) error{
	fmt.Printf("publish simple read to send message\n")
	// 申请队列
	r.Lock()
	defer r.Unlock()
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
		)
	if err != nil {
		loggo.ErrorFormat("declare queue failed,err:%v",err)
		return err
	}
	// 调用channel 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:[]byte(message),
		})
	fmt.Printf("publish simple send message err:%v\n",err)
	return r.failOnErr(err,"publish message to channel failed")
}

// step 三: 消费消息
func (r *RabbitMQ) ConsumerSimple(os service.OrderServer,ps service.GoodsServer) {
	// 申请队列
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
		)
	if err != nil {
		loggo.ErrorFormat("declare consumer channel faile,err:%v",err)
		return
	}
	// 配合Qos 消费者流控 防止暴库
	err = r.channel.Qos(
		// 每次消费消息数量 1
		1,
		// 服务器传递的最大容量 八位字节为单位
		0,
		// false 当前消息队列 true 全局使用
		false,
	)
	_ = r.failOnErr(err,"consumer channel Qos err")
	// 接收消息
	msgs,err := r.channel.Consume(
		// queue
		r.QueueName,
		// 区分多个消费者
		"",
		// 是否自动应答
		false,
		// 是否独有
		false,
		// 设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,
		// 队列是否阻塞 false为阻塞
		false,
		// 额外属性 args
		nil,
		)
	if err != nil {
		loggo.ErrorFormat("connect consumer channel failed,err:%v",err)
		return
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 消息逻辑处理
			loggo.InfoFormat("此处处理消息:%s",d.Body)
			message := &models.Message{}
			err := json.Unmarshal([]byte(d.Body),message)
			_ = r.failOnErr(err,"反序列化消息失败")
			//插入订单
			order := new(models.Order)
			order.UserId = message.UserID
			order.GoodsId = message.ProductID
			_,err = os.Insert(order)
			_ = r.failOnErr(err,"创建订单失败")
			//修改商品数据库
			pidStr := strconv.FormatInt(message.ProductID,10)
			goods,err := ps.SelectByField("id",pidStr)
			_ = r.failOnErr(err,"获取商品信息失败")
			goods.Num -= 1
			err = ps.Update(goods)
			_ = r.failOnErr(err,"更新商品数量失败")
			// 配合 autoAck 为false 使用 autoAck 为false 如果不写 d.Ack(false) 后果非常严重 会重新返回消息队列 等待消费
			// 如果d.Ack(true) 一般批量消息处理 表示确认所有未确认的消息
			// 如果d.Ack(false) 表示确认当前消息
			err = d.Ack(false)
			_ = r.failOnErr(err,"d.Ack err")
		}
	}()
	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C exist")
	<-forever
}

// 订阅模式实例
// 订阅模式 一个消息 可以同时被多个消费者消费
func NewRabbitMQSPubSub(exchange string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchange, "")
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	_ = rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	_ = rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}
// 订阅模式生产消息
func (r *RabbitMQ) PublishPub(message string) {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		// 交换机
		r.Exchange,
		// 交换机类型 广播
		"fanout",
		// 持久化
		true,
		// 自动删除
		false,
		// true标识此交换机不可以被client用来推送消息,仅用来进行 交换机 和 交换机 之间的绑定
		false,
		// 是否阻塞
		false,
		nil,
		)
	_ = r.failOnErr(err,"Declare an exchange failed")
	// 发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:[]byte(message),
		})
	_ = r.failOnErr(err,"Product message error")
	return
}
// 订阅模式消费消息
func (r *RabbitMQ) ConsumerSub() {
	// 1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", // 广播类型
		true,     // 持久化
		false,    // 是否删除
		false,    // true表示这个exchange不可以被client用来推送消息的，仅用来进行exchange和exchange之间的绑定
		false,
		nil,
	)
	_ = r.failOnErr(err, "Failed to declare a exchange")
	// 2. 试探性创建队列
	q, err := r.channel.QueueDeclare(
		"", // 随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	_ = r.failOnErr(err, "failed to declare a queue")
	// 绑定队列到 exchange中
	err = r.channel.QueueBind(
		q.Name,
		"", // 在订阅模式下，这里的key为空
		r.Exchange,
		false,
		nil)
	// 消费消息
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	// 3. 启用协程处理消息
	go func() {
		for d := range messages {
			// 实现我们要处理的逻辑函数
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages, to exit process CTRL+C")
	<-forever
}

// 路由模式 一个消息被多个消费者消费,并且,消息的目标队列 可以被 生产者指定
// 实现指定消息被 指定消费者消费
// 路由模式实例
func NewRabbitMQRouting(exchange,routingKey string) *RabbitMQ {
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchange, routingKey)
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	_ = rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	_ = rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}
// 路由模式生产消息
func (r *RabbitMQ) PublishRouting(message string) {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		// 指定为  direct模式
		"direct",
		true,
		false,
		false,
		false,
		nil,
		)
	_ = r.failOnErr(err,"declare exchange err")
	// 发送消息
	err = r.channel.Publish(
		r.Exchange,
		// 要设置key
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:[]byte(message),
		})
	_ = r.failOnErr(err,"Product message error")
	return
}

// 路由模式 消费
func (r *RabbitMQ) ConsumerRouting() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
		)
	_ = r.failOnErr(err,"declare an exchange err")
	q,err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,

		)
	_ = r.failOnErr(err,"declare queue err")
	// 绑定要交换机中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
		)
	// 消费消息
	// 接收消息
	msgs,err := r.channel.Consume(
		// queue
		q.Name,
		// 区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否独有
		false,
		// 设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,
		// 队列是否阻塞 false为阻塞
		false,
		// 额外属性 args
		nil,
	)
	if err != nil {
		loggo.ErrorFormat("connect consumer channel failed,err:%v",err)
		return
	}
	// 配合Qos 消费者流控 防止暴库
	err = r.channel.Qos(
		// 每次消费消息数量 1
		1,
		// 服务器传递的最大容量 八位字节为单位
		0,
		// false 当前消息队列 true 全局使用
		false,
	)
	_ = r.failOnErr(err,"consumer channel Qos err")
	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 消息逻辑处理
			loggo.InfoFormat("此处处理消息:%s",d.Body)
			// 配合 autoAck 为false 使用 autoAck 为false 如果不写 d.Ack(false) 后果非常严重 会重新返回消息队列 等待消费
			// 如果d.Ack(true) 一般批量消息处理 表示确认所有未确认的消息
			// 如果d.Ack(false) 表示确认当前消息
			err = d.Ack(false)
			_ = r.failOnErr(err,"d.Ack err")
		}
	}()
	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C exist")
	<-forever
}


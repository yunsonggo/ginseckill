
package rabbitmq
/*
import (
	"fmt"
	"strconv"
	"time"
)
// 简单模式生产
func ProductSimple() {
	imooc := NewRabbitMQSimple("imoocSimple")
	for i := 0; i < 100; i++ {
		imooc.PublishSimple("hello imooc" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
// 工作模式消费
func ConsumerSimple() {
	imooc := NewRabbitMQSimple("imoocSimple")
	imooc.ConsumerSimple()
}


// 工作模式生产
func ProductWork() {
	imooc := NewRabbitMQSimple("imoocWork")
	for i := 0; i < 100; i++ {
		imooc.PublishSimple("hello imooc" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

// 工作模式消费一
func ConsumerWorkOne() {
	imooc := NewRabbitMQSimple("imoocWork")
	imooc.ConsumerSimple()
}
// 工作模式消费二
func ConsumerWorkTwo() {
	imooc := NewRabbitMQSimple("imoocWork")
	imooc.ConsumerSimple()
}

//订阅模式生产
func ProduceHub() {
	imooc := NewRabbitMQSPubSub("exImooc")
	for i := 0; i < 100; i++ {
		imooc.PublishPub("订阅模式生产" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

// 订阅模式消费1
func ConsumerHubOne() {
	imoocOne := NewRabbitMQSPubSub("exImooc")
	imoocOne.ConsumerSub()
}
// 订阅模式消费2
func ConsumerHubTwo() {
	imoocTwo := NewRabbitMQSPubSub("exImooc")
	imoocTwo.ConsumerSub()
}


// routing 模式生产消息
func ProductRouting() {
	imoocOne := NewRabbitMQRouting("exImooc","imooc_one")
	imoocTwo := NewRabbitMQRouting("exImooc","imooc_two")

	for i := 0; i < 10; i++ {
		imoocOne.PublishRouting("hello imooc one!" + strconv.Itoa(i))
		imoocTwo.PublishRouting("hello imooc two!" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
// routing 消费端一
func ConsumerRoutingOne() {
	imoocOne := NewRabbitMQRouting("exImooc","imooc_one")
	imoocOne.ConsumerRouting()
}
// routing 消费端二
func ConsumerRoutingTwo() {
	imoocTwo := NewRabbitMQRouting("exImooc","imooc_two")
	imoocTwo.ConsumerRouting()
}
 */
package dial

import (
	"github.com/Shopify/sarama"
	"github.com/yunsonggo/loggo"
)

var KP sarama.SyncProducer

func KafkaProducer(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	cli , err := sarama.NewSyncProducer([]string{addr},config)
	if err != nil {
		loggo.Error("init kafka producer failed, err:", err)
		return
	}
	KP = cli
	return
}


func SendToKafka(data,topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid,offset,err := KP.SendMessage(msg)
	if err != nil {
		loggo.ErrorFormat("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}
	loggo.InfoFormat("send succ, pid:%v offset:%v, topic:%v\n", pid, offset, topic)
	return
}
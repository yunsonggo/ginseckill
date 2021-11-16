package dial

import (
	"github.com/Shopify/sarama"
	"github.com/yunsonggo/loggo"
	"strings"
	"sync"
)

type KafkaConsumerServer struct {
	Consumer sarama.Consumer
	Addr     string
	Topic    string
	Wg       sync.WaitGroup
}

var KCS *KafkaConsumerServer

func InitKfkCS(addr, topic string) (err error) {
	KCS = &KafkaConsumerServer{}
	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		loggo.ErrorFormat("init kafka consumer server failed, err:%v", err)
		return
	}
	KCS.Consumer = consumer
	KCS.Addr = addr
	KCS.Topic = topic
	loggo.Info("init kafka consumer success")
	return
}

/*
// 消费kafka 消息 到ES
func WatchMsg() {
	logs.Info("message consumer from kafka to es run")
	partitions, err := KFKCS.Consumer.Partitions(KFKCS.Topic)
	logs.Info("kafka consumer partitions:%v", partitions)
	if err != nil {
		logs.Error("kafka consumer server's partitions lost,err:", err)
		return
	}

	for partition := range partitions {
		pc, err := KFKCS.Consumer.ConsumePartition(KFKCS.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logs.Error("kafka consumer partition %d ,err : %s\n", partition, err)
			return
		}
		logs.Info("pc is a interface")
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			logs.Info("go pc ..")
			for msg := range pc.Messages() {
				logs.Info("send to es this topic: %s, msg: %s", KFKCS.Topic, msg.Value)
				err = SendMsgToES(KFKCS.Topic, msg.Value)
				if err != nil {
					logs.Warn("send to es failed, err:%v", err)
					continue
				}
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
}
 */
package dial

import (
	"context"
	elasticV7 "github.com/olivere/elastic/v7"
	"github.com/yunsonggo/loggo"
	"time"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

var ESClient *elasticV7.Client
var esAddr string

func InitESServer(addr string) (err error) {
	esAddr = addr
	cli, err := elasticV7.NewClient(elasticV7.SetSniff(false), elasticV7.SetURL(addr))
	if err != nil {
		loggo.Error("connect es server error", err)
		return
	}
	ESClient = cli
	loggo.Info("es client init success")
	return
}

func SendMsgToES(topic string, data []byte) (err error) {
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)
	info,code,err := ESClient.Ping(esAddr).Do(context.Background())
	time.Sleep(time.Millisecond * 2)
	if err != nil {
		loggo.Error("ping es server error:",err)
		return
	}
	loggo.InfoFormat("es info :%v,code :%d",info,code)
	exists,existsErr := ESClient.IndexExists(topic).Do(context.Background())
	time.Sleep(time.Millisecond * 2)
	if existsErr != nil {
		loggo.ErrorFormat("check es index:%s error:%v",topic,err)
		return
	}
	if !exists {
		createIndex,createErr := ESClient.CreateIndex(topic).BodyJson(msg).Do(context.Background())
		time.Sleep(time.Millisecond * 2)
		if createErr != nil {
			loggo.Error("create es index error")
			return
		}
		if !createIndex.Acknowledged {
			loggo.Error("create es index failed")
			return
		}
	}
	_, err = ESClient.Index().Index(topic).BodyJson(msg).Do(context.Background())
	if err != nil {
		loggo.ErrorFormat("send msg to es faild,err :%v", err)
		return
	}
	loggo.Info("send msg to es success")
	return
}

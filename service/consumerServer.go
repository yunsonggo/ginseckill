package service

import (
	"2022/ginseckill/dial"
	"2022/ginseckill/models"
)

type ConsumerServer interface {
	SelectByName(name string) (consumer *models.ConsumerModel,err error)
	InsertConsumer(consumer *models.ConsumerModel) (id int64,err error)
}

type consumerServer struct {}

func NewConsumerServer() ConsumerServer {
	return &consumerServer{}
}

func (cs *consumerServer) SelectByName(name string) (consumer *models.ConsumerModel,err error) {
	con:=new(models.ConsumerModel)
	con.Name = name
	_,err = dial.D.Where("name = ?",name).Get(con)
	consumer = con
	return
}
func (cs *consumerServer) InsertConsumer(consumer *models.ConsumerModel) (id int64,err error) {
	id,err = dial.D.InsertOne(consumer)
	return
}
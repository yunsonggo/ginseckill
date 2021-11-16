package dial

import (
	"2022/ginseckill/models"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var D *xorm.Engine

func MysqlEngine(conn,driver string,show bool) (err error) {
	db,err := xorm.NewEngine(driver,conn)
	if err != nil {
		panic(fmt.Sprintf("连接数据库异常:%v\n",err))
	}
	db.ShowSQL(show)
	err = db.Sync2(
		new(models.Goods),
		new(models.Order),
		new(models.ConsumerModel),
		)
	D = db
	return
}
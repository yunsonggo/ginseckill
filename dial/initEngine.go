package dial

import (
	"2022/ginseckill/config"
	"github.com/yunsonggo/loggo"
)

func InitEngines() (err error) {
	mc := config.Conf.Mysql
	rc := config.Conf.Redis
	ec := config.Conf.Etcd
	err = MysqlEngine(mc.MysqlAddr,mc.MysqlDriver,mc.MysqlShow)
	if err != nil {
		loggo.ErrorFormat("init mysql engine failed err:%v",err)
	}
	err = RedisClient(rc.RedisAddr,rc.RedisPwd,rc.RedisDb)
	if err != nil {
		loggo.ErrorFormat("init redis client failed err:%v",err)
	}
	err = EtcdClient(ec.EtcdAddr,ec.EtcdTimeout)
	if err != nil {
		loggo.ErrorFormat("init etcd client failed err:%v",err)
	}
	return
}

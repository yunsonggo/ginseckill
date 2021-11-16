package common

import (
	"2022/ginseckill/config"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/yunsonggo/loggo"
)

// 初始化
func InitSession (engine *gin.Engine) {
	var sessConf = config.Conf.Session
	var resConf = config.Conf.Redis
	store,err := redis.NewStore(sessConf.SessionSize,"tcp",resConf.RedisAddr,resConf.RedisPwd,[]byte(sessConf.SessionSecret))
	if err != nil {
		loggo.ErrorFormat("init session Store err:%v\n", err)
		panic(fmt.Sprintf("初始化sessionStore错误:%v\n", err))
	}
	if sessConf.SessionMaxage != 0 {
		store.Options(sessions.Options{
			MaxAge:   sessConf.SessionMaxage,
		})
	}
	engine.Use(sessions.Sessions(sessConf.SessionName,store))
}

// 设置session
func SetSession(ctx *gin.Context,key,value string) (err error) {
	session := sessions.Default(ctx)
	if session == nil {
		err := errors.New("ctx default session err")
		panic(fmt.Sprintf("%v\n",err))
	}
	session.Set(key,value)
	return session.Save()
}

// 获取session
func GetSession(ctx *gin.Context,key string) interface{} {
	session := sessions.Default(ctx)
	if session == nil {
		err := errors.New("ctx default session err")
		panic(fmt.Sprintf("%v\n",err))
	}
	return session.Get(key)
}
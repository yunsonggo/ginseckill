package dial

import redisV7 "github.com/go-redis/redis/v7"

var R *redisV7.Client

func RedisClient(addr,pass string,db int) (err error) {
	client := redisV7.NewClient(&redisV7.Options{
		Addr: addr,
		Password: pass,
		DB: db,
	})
	_, err = client.Ping().Result()
	return
}
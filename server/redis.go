package server

import (
	"github.com/garyburd/redigo/redis"
)

var redisConn redis.Conn

func RedisInit() {
	var err error
	redisConn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	if _, err := redisConn.Do("AUTH", "foobared"); err != nil {
		redisConn.Close()
		panic(err)
	}
}

func GetRedisConn() redis.Conn {
	return redisConn
}

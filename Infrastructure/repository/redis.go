package repository

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var (
	//Redisdb redis连接
	Redisdb redis.Conn
)

func init() {
	var err error
	Redisdb, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err)
	}
}

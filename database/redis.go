package database

import (
	"log"
	"test/config"

	"github.com/gomodule/redigo/redis"
)


func NewRedisConn(cfg config.Redis) redis.Conn {
	conn, err := redis.Dial("tcp", cfg.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

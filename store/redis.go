package store

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

type Redis struct {
	Conn    redis.Conn
	Counter int
}

func New(r redis.Conn) Redis {
	return Redis{Conn: r, Counter: 0}
}

func (redis *Redis) Insert(value string) {
	_, err := redis.Conn.Do("SET", "eldis" + strconv.Itoa(redis.Counter), value)
	if err != nil {
		log.Fatal(err)
	}

	redis.Counter++
}

func (redis *Redis) Fetch(id int) string {
	value, err := redis.Conn.Do("GET", "eldis" + strconv.Itoa(id))
	if err != nil {
		log.Print("This id doesn't exist")
	}

	return value
}

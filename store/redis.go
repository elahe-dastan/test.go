package store

import (
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

func (r *Redis) Insert(value string) {
	_, err := r.Conn.Do("SET", "eldis" + strconv.Itoa(r.Counter), value)
	if err != nil {
		log.Fatal(err)
	}

	r.Counter++
}

func (r *Redis) Fetch(id int) string {
	value, err := redis.String(r.Conn.Do("GET", "eldis" + strconv.Itoa(id)))
	if err != nil {
		log.Print("This id doesn't exist")
	}

	return value
}

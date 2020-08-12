package main

import (
	"fmt"
	"test/config"
	"test/database"
	"test/store"
)

func main() {
	cfg := config.Read("")
	redis := store.New(database.NewRedisConn(cfg.Redis))
	redis.Insert("parham")
	redis.Fetch(5)
	fmt.Println(redis.Fetch(0))
}
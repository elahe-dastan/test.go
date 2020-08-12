package main

import (
	"test/config"
	"test/database"
)

func main() {
	cfg := config.Read("")
	//redis := store.New(database.NewRedisConn(cfg.Redis))
	//redis.Insert("parham")
	//redis.Fetch(5)
	//fmt.Println(redis.Fetch(0))

	database.NewDBConn(cfg.Database)
}
package main

import (
	"fmt"
	"log"
	"test/config"
	"test/database"
	"test/model"
	"test/store"
)

func main() {
	cfg := config.Read("")
	//redis := store.New(database.NewRedisConn(cfg.Redis))
	//redis.Insert("parham")
	//redis.Fetch(5)
	//fmt.Println(redis.Fetch(0))

	d := store.NewDatabase(database.NewDBConn(cfg.Database))
	d.Create()

	p := model.Person{
		Name:   "raha",
		Number: "09336005978",
	}

	if err := d.Insert(p); err != nil {
		log.Println(err)
	}

	r, err := d.Retrieve("raha")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(r)
}
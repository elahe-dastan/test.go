package store

import (
	"database/sql"
	"log"
)

type Database struct {
	Conn *sql.DB
}

func (d Database) Create() {
	_, err := d.Conn.Exec("CREATE TABLE IF NOT EXISTS person (" +
		"id serial PRIMARY KEY," +
		"name VARCHAR," +
		"number VARCHAR" +
		");")
	if err != nil {
		log.Fatal(err)
	}
}


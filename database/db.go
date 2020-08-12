package database

import (
	"database/sql"
	"log"
	"test/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDBConn(cfg config.Database) *sql.DB {
	db, err := sql.Open("mysql", cfg.Cstring())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
package store

import (
	"database/sql"
	"errors"
	"log"
	"test/model"
)

var ErrNotFound = errors.New("person doesn't exist in the table")

type Database struct {
	Conn *sql.DB
}

func NewDatabase(c *sql.DB) Database{
	return Database{Conn: c}
}

func (d Database) Create() {
	_, err := d.Conn.Exec("CREATE TABLE IF NOT EXISTS person (" +
		"id serial PRIMARY KEY," +
		"name VARCHAR(250)," +
		"number VARCHAR(250)," +
		"UNIQUE (name)" +
		");")
	if err != nil {
		log.Fatal(err)
	}
}

func (d Database) Insert(p model.Person) error {
	_, err := d.Conn.Exec("INSERT INTO person (name, number) VALUES (?, ?)",
		p.Name, p.Number)

	return err
}

func (d Database) Retrieve(name string) (model.Person, error) {
	var p model.Person

	err := d.Conn.QueryRow("SELECT * FROM person WHERE name = ?;", name).Scan(
		&p.ID, &p.Name, &p.Number)

	if p.Name == "" {
		err = ErrNotFound
	}

	return p,err
}

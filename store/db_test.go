package store

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"test/model"
	"testing"
)

func TestDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	p := model.Person{
		Name:   "parham",
		Number: "09390909540",
	}

	mock.ExpectExec("INSERT INTO person").WithArgs(p.Name, p.Number).WillReturnResult(sqlmock.NewResult(1,1))
	rows := sqlmock.NewRows([]string{"id", "name", "number"})
	rows.AddRow(p.ID, p.Name, p.Number)
	mock.ExpectQuery("SELECT (.+) FROM person WHERE name").WithArgs(p.Name).WillReturnRows(rows)

	d := NewDatabase(db)

	err = d.Insert(p)

	assert.NoError(t, err)

	r, err := d.Retrieve("parham")

	assert.NoError(t, err)
	assert.Equal(t, p.Name, r.Name)
	assert.Equal(t, p.Number, r.Number)

}

package store

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"test/config"
	"test/database"
	"testing"
)

func TestRedis(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	r := NewRedis(database.NewRedisConn(config.Redis{Addr: s.Addr()}))

	r.Insert("foo")
	v := r.Fetch(0)

	assert.Equal(t, v, "foo")
}
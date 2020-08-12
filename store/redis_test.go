package store

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/alicebob/miniredis/v2"
)

func TestRedis(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	defer s.Close()

	err = s.Set("foo", "bar")

	assert.NoError(t, err)

	v, err := s.Get("foo")

	assert.NoError(t, err)
	assert.Equal(t, v, "bar")

}
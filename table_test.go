package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildInsert(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.insert.params = map[string]interface{}{
		"columnA": "value1",
	}
	query, args := s.BuildInsert()
	assert.Equal(t, "INSERT INTO users (columnA) VALUES (?);", query)
	assert.Equal(t, []interface{}{"value1"}, args)
}

func TestBuildQuery(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.limit = 10
	s.order = "DESC"
	s.orderKey = "email"
	s.where = struct {
		key   string
		value interface{}
	}{"name", "Taro"}
	assert.Equal(t, "SELECT id, name, email FROM users ORDER BY email DESC LIMIT 10 WHERE name = 'Taro';", s.BuildQuery())
}

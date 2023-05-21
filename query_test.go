package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildQuery(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.query.limit = 10
	s.query.order = "DESC"
	s.query.orderKey = "email"
	s.query.where = struct {
		key   string
		value interface{}
	}{"name", "Taro"}
	assert.Equal(t, "SELECT id, name, email FROM users WHERE name = 'Taro' ORDER BY email DESC LIMIT 10;", s.query.BuildQuery(s.columns, s.tableName))
}

func TestBuildInsert(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.query.insert.params = map[string]interface{}{
		"columnA": "value1",
	}
	query, args := s.query.BuildInsert(s.tableName)
	assert.Equal(t, "INSERT INTO users (columnA) VALUES (?);", query)
	assert.Equal(t, []interface{}{"value1"}, args)
}

func TestBuildDelete(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")

	query := s.query.BuildDelete(s.tableName, 12)
	assert.Equal(t, "DELETE FROM users WHERE id = 12;", query)
}

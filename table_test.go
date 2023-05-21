package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildInsert(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.query.insert.params = map[string]interface{}{
		"columnA": "value1",
	}
	query, args := s.BuildInsert()
	assert.Equal(t, "INSERT INTO users (columnA) VALUES (?);", query)
	assert.Equal(t, []interface{}{"value1"}, args)
}

func TestBuildDelete(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")

	query := s.BuildDelete(12)
	assert.Equal(t, "DELETE FROM users WHERE id = 12;", query)
}

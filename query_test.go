package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildQuery(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")

	t.Run("no condition", func(t *testing.T) {
		assert.Equal(t, "SELECT id, name, email FROM users;", s.query.BuildQuery(s.columns, s.tableName))
	})

	t.Run("where", func(t *testing.T) {
		s.query.where = struct {
			key   string
			value interface{}
		}{"name", "Taro"}
		assert.Equal(t, "SELECT id, name, email FROM users WHERE name = 'Taro';", s.query.BuildQuery(s.columns, s.tableName))

		// refresh query
		s.query = Query{}
	})

	t.Run("order", func(t *testing.T) {
		s.query.order = "DESC"
		s.query.orderKey = "email"
		assert.Equal(t, "SELECT id, name, email FROM users ORDER BY email DESC;", s.query.BuildQuery(s.columns, s.tableName))

		s.query = Query{}
	})

	t.Run("limit", func(t *testing.T) {
		s.query.limit = 10
		assert.Equal(t, "SELECT id, name, email FROM users LIMIT 10;", s.query.BuildQuery(s.columns, s.tableName))

		s.query = Query{}
	})

	t.Run("inner join", func(t *testing.T) {
		s.query.innerJoin.left = "users"
		s.query.innerJoin.right = "posts"
		s.query.innerJoin.hasMany = true
		assert.Equal(t, "SELECT id, name, email FROM users INNER JOIN posts on users.id = posts.user_id;", s.query.BuildQuery(s.columns, s.tableName))

		s.query = Query{}
	})

	t.Run("inner join", func(t *testing.T) {
		// change table name only this inner join case
		s.SetTable("posts")
		s.query.innerJoin.left = "posts"
		s.query.innerJoin.right = "users"
		s.query.innerJoin.hasMany = false
		assert.Equal(t, "SELECT id, name, email FROM posts INNER JOIN users on posts.user_id = users.id;", s.query.BuildQuery(s.columns, s.tableName))

		s.SetTable("users")
		s.query = Query{}
	})

	t.Run("where, order, limit", func(t *testing.T) {
		s.query.where = struct {
			key   string
			value interface{}
		}{"name", "Taro"}
		s.query.order = "DESC"
		s.query.orderKey = "email"
		s.query.limit = 10
		assert.Equal(t, "SELECT id, name, email FROM users WHERE name = 'Taro' ORDER BY email DESC LIMIT 10;", s.query.BuildQuery(s.columns, s.tableName))

		s.query = Query{}
	})
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

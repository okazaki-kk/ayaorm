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
		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users;", query)
		assert.Empty(t, args)
	})

	t.Run("where", func(t *testing.T) {
		s.query.Where("name", "Taro")

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE name = ?;", query)
		assert.Equal(t, []interface{}{"Taro"}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("where num", func(t *testing.T) {
		s.query.Where("age", 20)

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age = ?;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("where >", func(t *testing.T) {
		s.query.Where("age", ">", 20)

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ?;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("where null", func(t *testing.T) {
		s.query.Where("age", nil)

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age IS NULL;", query)
		assert.Empty(t, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("group by", func(t *testing.T) {
		s.query.groupBy = []string{"name", "email"}

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users GROUP BY name, email;", query)
		assert.Empty(t, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("group by and where", func(t *testing.T) {
		s.query.groupBy = []string{"name", "email"}
		s.query.Where("age", ">", 20)

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? GROUP BY name, email;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("having", func(t *testing.T) {
		s.query.groupBy = []string{"name", "email"}
		s.query.Having("COUNT(*)", ">", 1)

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users GROUP BY name, email HAVING COUNT(*) > ?;", query)
		assert.Equal(t, []interface{}{1}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("where and group by and having", func(t *testing.T) {
		s.query.Where("age", ">", 20)
		s.query.groupBy = []string{"name", "email"}
		s.query.Having("COUNT(*)", ">", 1)

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? GROUP BY name, email HAVING COUNT(*) > ?;", query)
		assert.Equal(t, []interface{}{20, 1}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("or", func(t *testing.T) {
		s.query.Where("age", ">", 20)
		s.query.Or("name", "Taro")

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? OR name = ?;", query)
		assert.Equal(t, []interface{}{20, "Taro"}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("or and null", func(t *testing.T) {
		s.query.Where("age", ">", 20)
		s.query.Or("name", nil)

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? OR name IS NULL;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.query = Query{}
	})

	t.Run("order", func(t *testing.T) {
		s.query.order = "DESC"
		s.query.orderKey = "email"

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users ORDER BY email DESC;", query)
		assert.Empty(t, args)

		s.query = Query{}
	})

	t.Run("limit", func(t *testing.T) {
		s.query.limit = 10
		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users LIMIT 10;", query)
		assert.Empty(t, args)

		s.query = Query{}
	})

	t.Run("inner join", func(t *testing.T) {
		s.query.innerJoin.left = "users"
		s.query.innerJoin.right = "posts"
		s.query.innerJoin.hasMany = true

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users INNER JOIN posts on users.id = posts.user_id;", query)
		assert.Empty(t, args)

		s.query = Query{}
	})

	t.Run("inner join", func(t *testing.T) {
		// change table name only this inner join case
		s.SetTable("posts")
		s.query.innerJoin.left = "posts"
		s.query.innerJoin.right = "users"
		s.query.innerJoin.hasMany = false

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM posts INNER JOIN users on posts.user_id = users.id;", query)
		assert.Empty(t, args)

		s.SetTable("users")
		s.query = Query{}
	})

	t.Run("where, order, limit", func(t *testing.T) {
		s.query.Where("name", "Taro")
		s.query.order = "DESC"
		s.query.orderKey = "email"
		s.query.limit = 10

		query, args := s.query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE name = ? ORDER BY email DESC LIMIT 10;", query)
		assert.Equal(t, []interface{}{"Taro"}, args)

		s.query = Query{}
	})
}

func TestBuildInsert(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.query.insert.params = map[string]interface{}{
		"name":  "name1",
		"value": "value2",
	}
	query, args := s.query.BuildInsert(s.tableName)
	assert.Equal(t, "INSERT INTO users (name, value) VALUES (?, ?);", query)
	assert.Equal(t, []interface{}{"name1", "value2"}, args)
}

func TestBuildCreateAll(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.query.createAll.params = map[string][]interface{}{
		"name":  {"name1", "name2", "name3"},
		"value": {"value1", "value2", "value3"},
	}
	query, args := s.query.BuildCreateAll(s.tableName)
	assert.Equal(t, "INSERT INTO users (name, value) VALUES (?, ?), (?, ?), (?, ?);", query)
	assert.Equal(t, []interface{}{"name1", "value1", "name2", "value2", "name3", "value3"}, args)
}

func TestBuildDelete(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")

	query := s.query.BuildDelete(s.tableName, 12)
	assert.Equal(t, "DELETE FROM users WHERE id = 12;", query)
}

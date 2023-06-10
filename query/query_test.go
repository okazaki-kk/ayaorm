package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildQuery(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")

	t.Run("no condition", func(t *testing.T) {
		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users;", query)
		assert.Empty(t, args)
	})

	t.Run("where", func(t *testing.T) {
		s.Where("name", "Taro")

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE name = ?;", query)
		assert.Equal(t, []interface{}{"Taro"}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("where num", func(t *testing.T) {
		s.Where("age", 20)

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age = ?;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("where >", func(t *testing.T) {
		s.Where("age", ">", 20)

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ?;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("where null", func(t *testing.T) {
		s.Query.Where("age", nil)

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age IS NULL;", query)
		assert.Empty(t, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("group by", func(t *testing.T) {
		s.Query.groupBy = []string{"name", "email"}

		query, args := s.Query.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users GROUP BY name, email;", query)
		assert.Empty(t, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("group by and where", func(t *testing.T) {
		s.groupBy = []string{"name", "email"}
		s.Where("age", ">", 20)

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? GROUP BY name, email;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("having", func(t *testing.T) {
		s.groupBy = []string{"name", "email"}
		s.Having("COUNT(*)", ">", 1)

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users GROUP BY name, email HAVING COUNT(*) > ?;", query)
		assert.Equal(t, []interface{}{1}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("where and group by and having", func(t *testing.T) {
		s.Where("age", ">", 20)
		s.groupBy = []string{"name", "email"}
		s.Having("COUNT(*)", ">", 1)

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? GROUP BY name, email HAVING COUNT(*) > ?;", query)
		assert.Equal(t, []interface{}{20, 1}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("or", func(t *testing.T) {
		s.Where("age", ">", 20)
		s.Or("name", "Taro")

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? OR name = ?;", query)
		assert.Equal(t, []interface{}{20, "Taro"}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("or and null", func(t *testing.T) {
		s.Where("age", ">", 20)
		s.Or("name", nil)

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE age > ? OR name IS NULL;", query)
		assert.Equal(t, []interface{}{20}, args)

		// refresh query
		s.Query = Query{}
	})

	t.Run("order", func(t *testing.T) {
		s.order = "DESC"
		s.orderKey = "email"

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users ORDER BY email DESC;", query)
		assert.Empty(t, args)

		s.Query = Query{}
	})

	t.Run("limit", func(t *testing.T) {
		s.limit = 10
		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users LIMIT 10;", query)
		assert.Empty(t, args)

		s.Query = Query{}
	})

	t.Run("inner join", func(t *testing.T) {
		s.innerJoin.left = "users"
		s.innerJoin.right = "posts"
		s.innerJoin.hasMany = true

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users INNER JOIN posts on users.id = posts.user_id;", query)
		assert.Empty(t, args)

		s.Query = Query{}
	})

	t.Run("inner join", func(t *testing.T) {
		// change table name only this inner join case
		s.SetTable("posts")
		s.innerJoin.left = "posts"
		s.innerJoin.right = "users"
		s.innerJoin.hasMany = false

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM posts INNER JOIN users on posts.user_id = users.id;", query)
		assert.Empty(t, args)

		s.SetTable("users")
		s.Query = Query{}
	})

	t.Run("where, order, limit", func(t *testing.T) {
		s.Where("name", "Taro")
		s.order = "DESC"
		s.orderKey = "email"
		s.limit = 10

		query, args := s.BuildQuery(s.columns, s.tableName)
		assert.Equal(t, "SELECT id, name, email FROM users WHERE name = ? ORDER BY email DESC LIMIT 10;", query)
		assert.Equal(t, []interface{}{"Taro"}, args)

		s.Query = Query{}
	})
}

func TestBuildInsert(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.insert.params = map[string]interface{}{
		"name":  "name1",
		"value": "value2",
	}
	query, args := s.BuildInsert(s.tableName)
	assert.Equal(t, "INSERT INTO users (name, value) VALUES (?, ?);", query)
	assert.Equal(t, []interface{}{"name1", "value2"}, args)
}

func TestBuildCreateAll(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")
	s.createAll.params = map[string][]interface{}{
		"name":  {"name1", "name2", "name3"},
		"value": {"value1", "value2", "value3"},
	}
	query, args := s.BuildCreateAll(s.tableName)
	assert.Equal(t, "INSERT INTO users (name, value) VALUES (?, ?), (?, ?), (?, ?);", query)
	assert.Equal(t, []interface{}{"name1", "value1", "name2", "value2", "name3", "value3"}, args)
}

func TestBuildDelete(t *testing.T) {
	s := &Table{}
	s.SetTable("users")
	s.SetColumns("id", "name", "email")

	query := s.BuildDelete(s.tableName, 12)
	assert.Equal(t, "DELETE FROM users WHERE id = 12;", query)
}

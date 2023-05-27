package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	db, _ = sql.Open("sqlite3", "./ayaorm.db")

	sqlStmt := `
		drop table if exists users;
		drop table if exists posts;
		drop table if exists comments;
		create table users (id integer primary key autoincrement, name text, age int, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')) );
		create table posts (id integer primary key autoincrement, content text, author text, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')) );
		create table comments (id integer primary key autoincrement, content text, author text, post_id integer, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')), foreign key (post_id) references posts(id) );
		CREATE TRIGGER trigger_test_updated_at_users AFTER UPDATE ON posts
		BEGIN
			UPDATE posts SET updated_at = DATETIME('now', 'localtime') WHERE rowid == NEW.rowid;
		END;
		CREATE TRIGGER trigger_test_updated_at_posts AFTER UPDATE ON posts
		BEGIN
			UPDATE posts SET updated_at = DATETIME('now', 'localtime') WHERE rowid == NEW.rowid;
		END;
		CREATE TRIGGER trigger_test_updated_at_comments AFTER UPDATE ON comments
		BEGIN
			UPDATE comments SET updated_at = DATETIME('now', 'localtime') WHERE rowid == NEW.rowid;
		END;
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}

	code := m.Run()
	defer db.Close()
	defer os.Remove("./ayaorm.db")
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	post, err := Post{}.Create(PostParams{Content: "Golang Post", Author: "Me"})
	assert.NoError(t, err)
	assert.Equal(t, "Golang Post", post.Content)
	assert.Equal(t, "Me", post.Author)
	assert.Equal(t, 1, post.Id)
	assert.Equal(t, post.CreatedAt, post.UpdatedAt)
	assert.Equal(t, 1, Post{}.Count())

	comment, err := Comment{}.Create(CommentParams{Content: "Fantastic", Author: "You", PostId: 1})
	assert.NoError(t, err)
	assert.Equal(t, "Fantastic", comment.Content)
	assert.Equal(t, "You", comment.Author)
	assert.Equal(t, 1, comment.PostId)
	assert.Equal(t, 1, comment.Id)
	assert.Equal(t, comment.CreatedAt, comment.UpdatedAt)
	assert.Equal(t, 1, Comment{}.Count())

	comment, err = Comment{}.Create(CommentParams{Content: "Bad", Author: "He", PostId: 1})
	assert.NoError(t, err)
	assert.Equal(t, "Bad", comment.Content)
	assert.Equal(t, "He", comment.Author)
	assert.Equal(t, 1, comment.PostId)
	assert.Equal(t, 2, comment.Id)
	assert.Equal(t, comment.CreatedAt, comment.UpdatedAt)
	assert.Equal(t, 2, Comment{}.Count())
}

func TestUpdate(t *testing.T) {
	post, err := Post{}.First()
	assert.NoError(t, err)

	err = post.Update(PostParams{Content: "Golang Post Updated", Author: "Me Updated"})
	assert.NoError(t, err)
	assert.Equal(t, "Golang Post Updated", post.Content)
	assert.Equal(t, "Me Updated", post.Author)
	assert.Equal(t, 1, post.Id)
}

func TestSave(t *testing.T) {
	post := Post{}
	post.Content = "Ruby Post"
	post.Author = "Matz"
	countBefore := Post{}.Count()

	err := post.Save()
	assert.NoError(t, err)
	assert.NotZero(t, post.Id)

	countAfter := Post{}.Count()
	assert.Equal(t, countBefore+1, countAfter)
}

func TestDelete(t *testing.T) {
	post, err := Post{}.Last()
	assert.NoError(t, err)
	countBefore := Post{}.Count()

	err = post.Delete()
	assert.NoError(t, err)

	countAfter := Post{}.Count()
	assert.Equal(t, countBefore-1, countAfter)
}

func TestWhere(t *testing.T) {
	posts, err := Post{}.Where("content", "Golang Post Updated").Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(posts))
}

func TestWhere1(t *testing.T) {
	posts, err := Post{}.Where("content", "Golang Post Updated").Where("author", "Me Updated").Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(posts))
}

func TestFind(t *testing.T) {
	post, err := Post{}.Find(1)
	assert.NoError(t, err)
	assert.Equal(t, "Golang Post Updated", post.Content)
	assert.Equal(t, "Me Updated", post.Author)
	assert.Equal(t, 1, post.Id)
}

func TestFindBy(t *testing.T) {
	post, err := Post{}.FindBy("content", "Golang Post Updated")
	assert.NoError(t, err)
	assert.Equal(t, "Golang Post Updated", post.Content)
	assert.Equal(t, "Me Updated", post.Author)
	assert.Equal(t, 1, post.Id)
}

func TestPluck(t *testing.T) {
	contents, err := Comment{}.Pluck("content")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(contents))
	assert.Equal(t, "Fantastic", contents[0])
	assert.Equal(t, "Bad", contents[1])
}

func TestOrder(t *testing.T) {
	_, err := User{}.Create(UserParams{Name: "Aya", Age: 20})
	assert.NoError(t, err)
	_, err = User{}.Create(UserParams{Name: "Yui", Age: 18})
	assert.NoError(t, err)

	users, err := User{}.Order("age", "desc").Query()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "Aya", users[0].Name)
	assert.Equal(t, 20, users[0].Age)
	assert.Equal(t, "Yui", users[1].Name)
	assert.Equal(t, 18, users[1].Age)
}

func TestWhere2(t *testing.T) {
	users, err := User{}.Where("age", 20).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Aya", users[0].Name)
	assert.Equal(t, 20, users[0].Age)

	users, err = User{}.Where("age", ">", 18).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Aya", users[0].Name)
	assert.Equal(t, 20, users[0].Age)

	users, err = User{}.Where("age", "<", 19).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Yui", users[0].Name)
	assert.Equal(t, 18, users[0].Age)

	users, err = User{}.Where("age", ">=", 18).Query()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))

	users, err = User{}.Where("age", ">=", 18).Limit(1).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))

	users, err = User{}.Where("age", ">=", 180).Query()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(users))
}

func TestHasMany(t *testing.T) {
	post, err := Post{}.First()
	assert.NoError(t, err)
	comments, err := post.Comments()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(comments))
	assert.Equal(t, "Fantastic", comments[0].Content)
	assert.Equal(t, "Bad", comments[1].Content)
}

func TestBelongsTo(t *testing.T) {
	comment, err := Comment{}.First()
	assert.NoError(t, err)
	post, err := comment.Post()
	assert.NoError(t, err)
	assert.Equal(t, "Golang Post Updated", post.Content)
	assert.Equal(t, "Me Updated", post.Author)
	assert.Equal(t, 1, post.Id)
}

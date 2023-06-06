package tests

import (
	"database/sql"
	"log"
	"os"
	"strings"
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
		create table users (id integer primary key autoincrement, name text not null, age int not null, address text, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')) );
		create table posts (id integer primary key autoincrement, content text not null, author text not null, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')) );
		create table comments (id integer primary key autoincrement, content text not null, author text not null, post_id integer not null, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')), foreign key (post_id) references posts(id) );
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

		insert into users (name, age, address) values ('Alice', 18, 'Tokyo');
		insert into users (name, age, address) values ('Bob', 20, 'Osaka');
		insert into users (name, age, address) values ('Carol', 22, 'Nagoya');
		insert into users (name, age, address) values ('Dave', 24, 'Fukuoka');
		insert into users (name, age, address) values ('Eve', 26, 'Sapporo');
		insert into users (name, age, address) values ('Frank', 28, 'Okinawa');

		insert into posts (content, author) values ('Golang Post', 'Me');
		insert into posts (content, author) values ('Ruby Post', 'You');
		insert into posts (content, author) values ('Python Post', 'He');
		insert into posts (content, author) values ('Java Post', 'She');
		insert into posts (content, author) values ('C++ Post', 'They');
		insert into posts (content, author) values ('Ruby Post', 'We');
		insert into posts (content, author) values ('PHP Post', 'Us');

		insert into comments (content, author, post_id) values ('Fantastic', 'You', 1);
		insert into comments (content, author, post_id) values ('Great', 'He', 1);
		insert into comments (content, author, post_id) values ('Good', 'She', 2);
		insert into comments (content, author, post_id) values ('Bad', 'They', 3);
		insert into comments (content, author, post_id) values ('Terrible', 'We', 3);
		insert into comments (content, author, post_id) values ('Awful', 'Us', 4);
		insert into comments (content, author, post_id) values ('Horrible', 'You', 5);
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
	post, err := Post{}.Create(PostParams{Content: "Fortran Post", Author: "Me"})
	assert.NoError(t, err)
	assert.Equal(t, "Fortran Post", post.Content)
	assert.Equal(t, "Me", post.Author)
	assert.NotZero(t, post.Id)
	assert.Equal(t, post.CreatedAt, post.UpdatedAt)
	assert.Equal(t, 8, Post{}.Count())

	comment, err := Comment{}.Create(CommentParams{Content: "Oh My God", Author: "It", PostId: 1})
	assert.NoError(t, err)
	assert.Equal(t, "Oh My God", comment.Content)
	assert.Equal(t, "It", comment.Author)
	assert.Equal(t, 1, comment.PostId)
	assert.NotZero(t, comment.Id)
	assert.Equal(t, comment.CreatedAt, comment.UpdatedAt)
	assert.Equal(t, 8, Comment{}.Count())
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
	post.Content = "Rails Post"
	post.Author = "DHH"
	countBefore := Post{}.Count()

	err := post.Save()
	assert.NoError(t, err)
	assert.NotZero(t, post.Id)

	countAfter := Post{}.Count()
	assert.Equal(t, countBefore+1, countAfter)

	lastPost, err := Post{}.Last()
	assert.NoError(t, err)
	assert.Equal(t, "Rails Post", lastPost.Content)
	assert.Equal(t, "DHH", lastPost.Author)
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
	posts, err := Post{}.Where("content", "Ruby Post").Query()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(posts))
	assert.Equal(t, "Ruby Post", posts[0].Content)
	assert.Equal(t, "Ruby Post", posts[1].Content)
}

func TestWhere1(t *testing.T) {
	posts, err := Post{}.Where("content", "C++ Post").Where("author", "They").Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(posts))
	assert.Equal(t, "C++ Post", posts[0].Content)
	assert.Equal(t, "They", posts[0].Author)
}

func TestFind(t *testing.T) {
	post, err := Post{}.Find(2)
	assert.NoError(t, err)
	assert.Equal(t, "Ruby Post", post.Content)
	assert.Equal(t, "You", post.Author)
	assert.Equal(t, 2, post.Id)
}

func TestFindBy(t *testing.T) {
	post, err := Post{}.FindBy("content", "Ruby Post")
	assert.NoError(t, err)
	assert.Equal(t, "Ruby Post", post.Content)
	assert.Equal(t, "You", post.Author)
}

func TestPluck(t *testing.T) {
	contents, err := Comment{}.Pluck("content")
	assert.NoError(t, err)
	assert.Equal(t, 8, len(contents))
	assert.Equal(t, []interface{}{"Fantastic", "Great", "Good", "Bad", "Terrible", "Awful", "Horrible", "Oh My God"}, contents)
}

func TestOrder(t *testing.T) {
	users, err := User{}.Order("age", "desc").Query()
	assert.NoError(t, err)
	assert.Equal(t, 6, len(users))
	assert.Equal(t, 28, users[0].Age)
	assert.Equal(t, 26, users[1].Age)
	assert.Equal(t, 24, users[2].Age)
	assert.Equal(t, 22, users[3].Age)
	assert.Equal(t, 20, users[4].Age)
	assert.Equal(t, 18, users[5].Age)
}

func TestWhere2(t *testing.T) {
	users, err := User{}.Where("age", 20).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Bob", users[0].Name)
	assert.Equal(t, 20, users[0].Age)

	users, err = User{}.Where("age", ">", 27).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Frank", users[0].Name)
	assert.Equal(t, 28, users[0].Age)

	users, err = User{}.Where("age", "<", 19).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, 18, users[0].Age)

	users, err = User{}.Where("age", ">=", 26).Query()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))

	users, err = User{}.Where("age", ">=", 18).Limit(3).Query()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(users))

	users, err = User{}.Where("age", ">=", 180).Query()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(users))
}

func TestOr(t *testing.T) {
	users, err := User{}.Where("age", 20).Or("name", "Alice").Query()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, 18, users[0].Age)
	assert.Equal(t, "Bob", users[1].Name)
	assert.Equal(t, 20, users[1].Age)
}

func TestGroupBy(t *testing.T) {
	posts, err := Post{}.GroupBy("author").Query()
	assert.NoError(t, err)
	assert.Equal(t, 8, len(posts))
	assert.Equal(t, "He", posts[0].Author)

	_, err = Post{}.Create(PostParams{Content: "Golang Post", Author: "He"})
	assert.NoError(t, err)

	posts, err = Post{}.GroupBy("author").Query()
	assert.NoError(t, err)
	assert.Equal(t, 8, len(posts))
}

func TestHaving(t *testing.T) {
	posts, err := Post{}.GroupBy("author").Having("count(*)", 2).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(posts))
	assert.Equal(t, "He", posts[0].Author)
}

func TestNull(t *testing.T) {
	user, err := User{}.Create(UserParams{Name: "Null Name", Age: 50})
	assert.NoError(t, err)
	assert.Equal(t, "Null Name", user.Name)
	assert.Equal(t, 50, user.Age)
	assert.False(t, user.Address.Valid())

	users, err := User{}.Where("address", nil).Query()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Null Name", users[0].Name)
	assert.Equal(t, 50, users[0].Age)
	assert.False(t, users[0].Address.Valid())

	users, err = User{}.Where("name", "Eve").Or("address", nil).Query()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
}

func TestHasMany(t *testing.T) {
	post, err := Post{}.First()
	assert.NoError(t, err)
	comments, err := post.Comments()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(comments))
}

func TestBelongsTo(t *testing.T) {
	comment, err := Comment{}.Find(3)
	assert.NoError(t, err)
	post, err := comment.Post()
	assert.NoError(t, err)
	assert.Equal(t, "Ruby Post", post.Content)
	assert.Equal(t, "You", post.Author)
	assert.Equal(t, 2, post.Id)
}

func TestIsValid(t *testing.T) {
	t.Run("presence valid", func(t *testing.T) {
		post := Post{}
		post.Content = "Ruby Post"
		post.Author = "Matz"
		valid, err := post.IsValid()
		assert.Equal(t, 0, len(err))
		assert.True(t, valid)
	})

	t.Run("presence invalid", func(t *testing.T) {
		post := Post{}
		post.Content = "Ruby Post"
		valid, err := post.IsValid()
		assert.Equal(t, 1, len(err))
		assert.Equal(t, "author can't be blank", err[0].Error())
		assert.False(t, valid)
	})

	t.Run("length valid", func(t *testing.T) {
		post := Post{}
		post.Content = "Ruby Post"
		post.Author = "Matz"
		valid, err := post.IsValid()
		assert.Equal(t, 0, len(err))
		assert.True(t, valid)
	})

	t.Run("length invalid", func(t *testing.T) {
		post := Post{}
		post.Content = "Ruby Post Post Post Post Post Post"
		post.Author = "Matz"

		valid, err := post.IsValid()
		assert.Equal(t, 1, len(err))
		assert.Equal(t, "content is too long (maximum is 20 characters)", err[0].Error())
		assert.False(t, valid)

		post.Content = "Ru"
		valid, err = post.IsValid()
		assert.Equal(t, 1, len(err))
		assert.Equal(t, "content is too short (minimum is 3 characters)", err[0].Error())
		assert.False(t, valid)
	})

	t.Run("numericality valid", func(t *testing.T) {
		user := User{}
		user.Name = "Aya"
		user.Age = 20

		valid, err := user.IsValid()
		assert.Equal(t, 0, len(err))
		assert.True(t, valid)
	})

	t.Run("numericality invalid", func(t *testing.T) {
		user := User{}
		user.Name = "Aya"
		user.Age = -1

		valid, err := user.IsValid()
		assert.Equal(t, 1, len(err))
		assert.Equal(t, "age must be positive", err[0].Error())
		assert.False(t, valid)
	})

	t.Run("custom invalid", func(t *testing.T) {
		user := User{}
		user.Name = "custom-example"
		user.Age = 20

		valid, err := user.IsValid()
		assert.Equal(t, 1, len(err))
		assert.Equal(t, "name must not be custom-example", err[0].Error())
		assert.False(t, valid)
	})
}

func TestValidation(t *testing.T) {
	t.Run("validation before create", func(t *testing.T) {
		post, errs := Post{}.Create(PostParams{Content: "Golang Post Updated Updated Updated Updated", Author: "He"})
		assert.Equal(t, "content is too long (maximum is 20 characters)", errs.Error())
		assert.Empty(t, post)
	})

	t.Run("validation before update", func(t *testing.T) {
		post, err := Post{}.Find(1)
		assert.NoError(t, err)

		errs := post.Update(PostParams{Content: "Golang Post Updated Updated Updated Updated"})
		assert.Equal(t, "content is too long (maximum is 20 characters)", errs.Error())
	})

	t.Run("validation before save", func(t *testing.T) {
		post, err := Post{}.Find(1)
		assert.NoError(t, err)

		post.Content = "Golang Post Updated Updated Updated Updated"
		errs := post.Save()
		assert.Equal(t, "content is too long (maximum is 20 characters)", errs.Error())
	})

	t.Run("validate multi errors", func(t *testing.T) {
		post := Post{}
		post.Content = "Golang Post Updated Updated Updated Updated"

		err := post.Save()
		assert.True(t, strings.Contains(err.Error(), "author can't be blank"))
		assert.True(t, strings.Contains(err.Error(), "content is too long (maximum is 20 characters)"))
	})
}

func TestDeleteDependent(t *testing.T) {
	post, err := Post{}.Find(3)
	assert.NoError(t, err)

	comments, err := post.Comments()
	assert.NoError(t, err)
	beforeCount := Comment{}.Count()

	err = post.DeleteDependent()
	assert.NoError(t, err)

	afterCount := Comment{}.Count()
	assert.Equal(t, beforeCount-len(comments), afterCount)
}

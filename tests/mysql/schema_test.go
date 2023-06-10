package test_mysql

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=true")

	sqlStmt := `drop table if exists users;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}
	sqlStmt = `drop table if exists comments;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}
	sqlStmt = `drop table if exists posts;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}
	sqlStmt = `drop table if exists projects;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}

	sqlStmt = `create table posts (id int not null primary key auto_increment, content text not null, author text not null, created_at datetime default current_timestamp, updated_at datetime default current_timestamp );`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}

	sqlStmt = `create table users (id int not null primary key auto_increment, name text not null, age int not null, age1 int not null, age2 int not null, address text, created_at datetime default current_timestamp, updated_at datetime default current_timestamp );`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}

	sqlStmt = `create table comments (id int not null primary key auto_increment, content text not null, author text not null, post_id int not null, created_at datetime default current_timestamp, updated_at datetime default current_timestamp );`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}

	sqlStmt = `create table projects (id int not null primary key auto_increment, name text not null, post_id int not null, created_at datetime default current_timestamp, updated_at datetime default current_timestamp );`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}

	sqlStmt = `insert into users (name, age, age1, age2, address) values ('Alice', 18, 100, 100, 'Tokyo'), ('Bob', 20, 100, 100, 'Osaka'), ('Carol', 22, 100, 100, 'Nagoya'),  ('Dave', 24, 100, 100, 'Fukuoka'), ('Eve', 26, 100, 100, 'Sapporo'), ('Frank', 28, 100, 100, 'Okinawa');`
	_, err = db.Exec(strings.Replace(sqlStmt, "\n", "", -1))
	if err != nil {
		log.Fatal("db create error", err)
	}

	sqlStmt  = `insert into posts (content, author) values ('Golang Post', 'Me'), ('Ruby Post', 'You'), ('Python Post', 'He'), ('Java Post', 'She'), ('C++ Post', 'They'), ('Ruby Post', 'We'), ('PHP Post', 'Us');`
	_, err = db.Exec(strings.Replace(sqlStmt, "\n", "", -1))
	if err != nil {
		log.Fatal("db create error", err)
	}

	sqlStmt = `insert into comments (content, author, post_id) values ('Fantastic', 'You', 1), ('Great', 'He', 1), ('Good', 'She', 2), ('Bad', 'They', 3), ('Terrible', 'We', 3), ('Awful', 'Us', 4), ('Horrible', 'You', 5);`
	_, err = db.Exec(strings.Replace(sqlStmt, "\n", "", -1))
	if err != nil {
		log.Fatal("db create error", err)
	}

	code := m.Run()
	defer db.Close()
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

/*func TestPluck(t *testing.T) {
	contents, err := Comment{}.Pluck("content")
	assert.NoError(t, err)
	assert.Equal(t, 8, len(contents))
	assert.Equal(t, []interface{}{"Fantastic", "Great", "Good", "Bad", "Terrible", "Awful", "Horrible", "Oh My God"}, contents)
}*/

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

/*func TestGroupBy(t *testing.T) {
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
}*/

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

func TestJoin1(t *testing.T) {
	comments, err := Comment{}.JoinPost().Query()
	assert.NoError(t, err)
	assert.Equal(t, 8, len(comments))
	assert.Equal(t, 1, comments[0].PostId)
	assert.Equal(t, "Fantastic", comments[0].Content)
	assert.Equal(t, "You", comments[0].Author)
	assert.Equal(t, 1, comments[1].PostId)
	assert.Equal(t, "Great", comments[1].Content)
	assert.Equal(t, "He", comments[1].Author)
	assert.Equal(t, 2, comments[2].PostId)
	assert.Equal(t, "Good", comments[2].Content)
	assert.Equal(t, "She", comments[2].Author)
}

func TestJoin2(t *testing.T) {
	posts, err := Post{}.JoinComments().Query()
	assert.NoError(t, err)
	assert.Equal(t, 8, len(posts))
	assert.Equal(t, 2, posts[2].Id)
	assert.Equal(t, "Ruby Post", posts[2].Content)
	assert.Equal(t, "You", posts[2].Author)
}

/*func TestHasOne(t *testing.T) {
	post, err := Post{}.Last()
	assert.NoError(t, err)
	_, err = Project{}.Create(ProjectParams{Name: "Project-Post", PostId: post.Id})
	assert.NoError(t, err)

	porject, err := post.Project()
	assert.NoError(t, err)
	assert.Equal(t, "Project-Post", porject.Name)
}*/

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

	t.Run("onCreate validation", func(t *testing.T) {
		_, err := User{}.Create(UserParams{Name: "custom-example1", Age1: -20})
		assert.Equal(t, "age1 must be positive", err.Error())

		u, err := User{}.Last()
		assert.NoError(t, err)
		err = u.Update(UserParams{Name: "custom-example1", Age1: -20})
		assert.NoError(t, err)
	})

	t.Run("onUpdate validation", func(t *testing.T) {
		_, err := User{}.Create(UserParams{Name: "custom-example1", Age2: -20})
		assert.NoError(t, err)

		u, err := User{}.Last()
		assert.NoError(t, err)
		err = u.Update(UserParams{Name: "custom-example1", Age1: -20})
		assert.Equal(t, "age2 must be positive", err.Error())
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

func TestCreateAll(t *testing.T) {
	posts := []PostParams{
		{Content: "Golang Post", Author: "Me"},
		{Content: "Ruby Post", Author: "You"},
		{Content: "Python Post", Author: "He"},
		{Content: "Java Post", Author: "She"},
		{Content: "C++ Post", Author: "They"},
		{Content: "Ruby Post", Author: "We"},
		{Content: "PHP Post", Author: "Us"},
	}

	err := Post{}.CreateAll(posts)
	assert.NoError(t, err)
	assert.Equal(t, 14, Post{}.Count())
}

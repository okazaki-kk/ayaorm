package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ = sql.Open("sqlite3", "./ayaorm.db")

	sqlStmt := `
		drop table if exists posts;
		drop table if exists comments;
		create table posts (id integer primary key autoincrement, content text, author text, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')) );
		create table comments (id integer primary key autoincrement, content text, author text, post_id integer, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')), foreign key (post_id) references posts(id) );
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("db create error", err)
	}

	_, err = Post{}.Create(PostParams{Content: "Golang Post", Author: "Me"})
	if err != nil {
		log.Fatal("Post Create Error", err)
	}

	_, err = Comment{}.Create(CommentParams{Content: "Fantastic", Author: "You", PostId: 1})
	if err != nil {
		log.Fatal("Comment Create Error", err)
	}
	_, err = Comment{}.Create(CommentParams{Content: "Bad", Author: "He", PostId: 1})
	if err != nil {
		log.Fatal("Comment Create Error", err)
	}

	post, _ := Post{}.First()
	comments, err := post.Comments()
	if err != nil {
		log.Fatal("post.Comments.Error", err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}

	posts, err := Post{}.JoinComments().Query()

	if err != nil {
		log.Fatal("joinComment.Error", err)
	}
	for _, post := range posts {
		fmt.Println(post, "post")
	}
}

func (c Comment) String() string {
	return fmt.Sprintf("{ID: %d, Author: %s, Content: %s, CreatedAt: %s, UpdatedAt: %s}", c.Id, c.Author, c.Content, c.CreatedAt.Format("2006/01/02 15:04:05.000"), c.UpdatedAt.Format("2006/01/02 15:04:05.000"))
}

func (c Post) String() string {
	return fmt.Sprintf("{ID: %d, Author: %s, Content: %s, CreatedAt: %s, UpdatedAt: %s}", c.Id, c.Author, c.Content, c.CreatedAt.Format("2006/01/02 15:04:05.000"), c.UpdatedAt.Format("2006/01/02 15:04:05.000"))
}

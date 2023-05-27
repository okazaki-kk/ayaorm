package tests

import "github.com/okazaki-kk/ayaorm"

type Comment struct {
	ayaorm.Schema
	Content string
	Author  string
	PostId  int
}

type Post struct {
	ayaorm.Schema
	Content string
	Author  string
}

type User struct {
	ayaorm.Schema
	Name string
	Age  int
}

func (m Post) hasManyComments() {
}

func (m Comment) belongsToPost() {
}

func (m Post) validatesPresenceOfAuthor() ayaorm.Rule {
	return ayaorm.MakeRule().Presence()
}

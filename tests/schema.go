package tests

import (
	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/null"
)

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
	Name    string
	Age     int
	Address null.NullString
}

func (m Post) hasManyComments() {
}

func (m Comment) belongsToPost() {
}

func (m Post) validatesPresenceOfAuthor() ayaorm.Rule {
	return ayaorm.MakeRule().Presence()
}

func (m Post) validateLengthOfContent() ayaorm.Rule {
	return ayaorm.MakeRule().MaxLength(10).MinLength(3)
}

func (m User) validateNumericalityOfAge() ayaorm.Rule {
	return ayaorm.MakeRule().Numericality().OnlyInteger()
}

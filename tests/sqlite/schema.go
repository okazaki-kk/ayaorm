package test_sqlite

import (
	"errors"

	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/null"
	"github.com/okazaki-kk/ayaorm/validate"
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

type Project struct {
	ayaorm.Schema
	Name   string
	PostId int
}

type User struct {
	ayaorm.Schema
	Name    string
	Age     int
	Age1    int
	Age2    int
	Address null.NullString
}

func (m Post) hasManyComments() {
}

func (m Comment) belongsToPost() {
}

func (m Project) belongsToPost() {
}

func (m Post) hasOneProject() {
}

func (m Post) validatesPresenceOfAuthor() validate.Rule {
	return validate.MakeRule().Presence()
}

func (m Post) validateLengthOfContent() validate.Rule {
	return validate.MakeRule().MaxLength(20).MinLength(3)
}

func (m User) validateNumericalityOfAge() validate.Rule {
	return validate.MakeRule().Numericality().Positive()
}

func (m User) validateNumericalityOfAge1() validate.Rule {
	return validate.MakeRule().Numericality().Positive().OnCreate()
}

func (m User) validateNumericalityOfAge2() validate.Rule {
	return validate.MakeRule().Numericality().Positive().OnUpdate()
}

func (m User) validateCustomRule() validate.Rule {
	return validate.CustomRule(func(es *[]error) {
		if m.Name == "custom-example" {
			*es = append(*es, errors.New("name must not be custom-example"))
		}
	})
}

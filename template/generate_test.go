package template

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	var fileInspect = FileInspect{
		PackageName: "main",
		StructInspect: []StructInspect{
			{
				ModelName:   "User",
				FieldKeys:   []string{"Id", "Name", "Age", "CreatedAt", "UpdatedAt"},
				FieldValues: []string{"int", "string", "int", "time.Time", "time.Time"},
			},
		},
		FuncInspect: []FuncInspect{
			{
				FuncName: "hasManyPosts",
				Recv:     "Comment",
			},
		},
	}

	err := Generate("user.go", fileInspect)
	assert.NoError(t, err)

	filePath := strings.ToLower("User") + "_gen.go"
	defer os.Remove(filePath)
	defer os.Remove("db_gen.go")

	data, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedTextBody, string(data))
}

func TestGenerateDB(t *testing.T) {
	err := GenerateDB()
	assert.NoError(t, err)

	defer os.Remove("db_gen.go")

	data, err := os.ReadFile("db_gen.go")
	assert.NoError(t, err)
	assert.Equal(t, expectedDBTextBody, string(data))
}

var expectedTextBody = `// Code generated by ayaorm. DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/okazaki-kk/ayaorm"
)

type UserRelation struct {
	model *User
	*ayaorm.Relation
}

func (m *User) newRelation() *UserRelation {
	r := &UserRelation{
		m,
		ayaorm.NewRelation(db).SetTable("users"),
	}
	r.Select(
		"id",
		"name",
		"age",
		"created_at",
		"updated_at",
	)

	return r
}

func (m User) Select(columns ...string) *UserRelation {
	return m.newRelation().Select(columns...)
}

func (r *UserRelation) Select(columns ...string) *UserRelation {
	cs := []string{}
	for _, c := range columns {
		if r.model.isColumnName(c) {
			cs = append(cs, fmt.Sprintf("users.%s", c))
		} else {
			cs = append(cs, c)
		}
	}
	r.Relation.SetColumns(cs...)
	return r
}

type UserParams User

func (m User) Build(p UserParams) *User {
	return &User{
		Schema: ayaorm.Schema{Id: p.Id},
		Name:   p.Name,
		Age:    p.Age,
	}
}

func (u User) Create(params UserParams) (*User, error) {
	user := u.Build(params)
	return u.newRelation().Create(user)
}

func (r *UserRelation) Create(user *User) (*User, error) {
	err := user.Save()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Update(params UserParams) error {
	return u.newRelation().Update(u.Id, params)
}

func (r *UserRelation) Update(id int, params UserParams) error {
	fieldMap := make(map[string]interface{})
	for _, c := range r.Relation.GetColumns() {
		switch c {
		case "name", "users.name":
			fieldMap["name"] = r.model.Name
		case "age", "users.age":
			fieldMap["age"] = r.model.Age
		}
	}
	return r.Relation.Update(id, fieldMap)
}

func (m *User) Save() error {
	lastId, err := m.newRelation().Save()
	if m.Id == 0 {
		m.Id = lastId
	}
	return err
}

func (r *UserRelation) Save() (int, error) {
	fieldMap := make(map[string]interface{})
	for _, c := range r.Relation.GetColumns() {
		switch c {
		case "name", "users.name":
			fieldMap["name"] = r.model.Name
		case "age", "users.age":
			fieldMap["age"] = r.model.Age
		}
	}

	return r.Relation.Save(fieldMap)
}

func (m *User) Delete() error {
	return m.newRelation().Delete(m.Id)
}

func (m User) Count(column ...string) int {
	return m.newRelation().Count(column...)
}

func (m User) All() *UserRelation {
	return m.newRelation()
}

func (m User) Limit(limit int) *UserRelation {
	return m.newRelation().Limit(limit)
}

func (r *UserRelation) Limit(limit int) *UserRelation {
	r.Relation.Limit(limit)
	return r
}

func (m User) Order(key, order string) *UserRelation {
	return m.newRelation().Order(key, order)
}

func (r *UserRelation) Order(key, order string) *UserRelation {
	r.Relation.Order(key, order)
	return r
}

func (m User) Where(column string, value interface{}) *UserRelation {
	return m.newRelation().Where(column, value)
}

func (r *UserRelation) Where(column string, value interface{}) *UserRelation {
	r.Relation.Where(column, value)
	return r
}

func (m User) First() (*User, error) {
	return m.newRelation().First()
}

func (r *UserRelation) First() (*User, error) {
	r.Relation.First()
	return r.QueryRow()
}

func (m User) Last() (*User, error) {
	return m.newRelation().Last()
}

func (r *UserRelation) Last() (*User, error) {
	r.Relation.Last()
	return r.QueryRow()
}

func (m User) Find(id int) (*User, error) {
	return m.newRelation().Find(id)
}

func (r *UserRelation) Find(id int) (*User, error) {
	r.Relation.Find(id)
	return r.QueryRow()
}

func (m User) FindBy(column string, value interface{}) (*User, error) {
	return m.newRelation().FindBy(column, value)
}

func (r *UserRelation) FindBy(column string, value interface{}) (*User, error) {
	r.Relation.FindBy(column, value)
	return r.QueryRow()
}

func (m User) Pluck(column string) ([]interface{}, error) {
	return m.newRelation().Pluck(column)
}

func (r *UserRelation) Pluck(column string) ([]interface{}, error) {
	return r.Relation.Pluck(column)
}

func (r *UserRelation) Query() ([]*User, error) {
	rows, err := r.Relation.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*User{}
	for rows.Next() {
		row := &User{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *UserRelation) QueryRow() (*User, error) {
	row := &User{}
	err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *User) fieldPtrByName(name string) interface{} {
	switch name {
	case "id", "users.id":
		return &m.Id
	case "name", "users.name":
		return &m.Name
	case "age", "users.age":
		return &m.Age
	case "created_at", "users.created_at":
		return &m.CreatedAt
	case "updated_at", "users.updated_at":
		return &m.UpdatedAt
	default:
		return nil
	}
}

func (m *User) fieldPtrsByName(names []string) []interface{} {
	fields := []interface{}{}
	for _, n := range names {
		f := m.fieldPtrByName(n)
		fields = append(fields, f)
	}
	return fields
}

func (m *User) isColumnName(name string) bool {
	for _, c := range m.columnNames() {
		if c == name {
			return true
		}
	}
	return false
}

func (m *User) columnNames() []string {
	return []string{
		"id",
		"name",
		"age",
		"created_at",
		"updated_at",
	}
}

func (m Post) Comments() ([]*Comment, error) {
	c, err := Comment{}.Where("post_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (u Post) JoinComments() *PostRelation {
	return u.newRelation().JoinComments()
}

func (u *PostRelation) JoinComments() *PostRelation {
	u.Relation.InnerJoin("posts", "comments")
	return u
}
`

var expectedDBTextBody = `// Code generated by ayaorm. DO NOT EDIT.
package main

import "database/sql"

var db *sql.DB
`

package template

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	var fileInspect = FileInspect{
		PackageName: "testss",
		StructInspect: []StructInspect{
			{
				ModelName:   "Comment",
				FieldKeys:   []string{"Id", "Content", "Author", "PostId", "CreatedAt", "UpdatedAt"},
				FieldValues: []string{"int", "string", "string", "int", "time.Time", "time.Time"},
			},
		},
		FuncInspect: []FuncInspect{
			{
				FuncName: "hasManyComments",
				Recv:     "Post",
				HasMany:  true,
			},
			{
				FuncName: "belongsToPost",
				Recv:     "Comment",
				BelongTo: true,
			},
			{
				FuncName:         "validatesPresenceOfAuthor",
				Recv:             "Post",
				ValidatePresence: true,
			},
			{
				FuncName:       "validateLengthOfContent",
				Recv:           "Post",
				ValidateLength: true,
			},
			{
				FuncName:             "validateNumericalityOfAge",
				Recv:                 "User",
				ValidateNumericality: true,
			},
		},
	}

	err := Generate("schema.go", fileInspect)
	assert.NoError(t, err)

	filePath := strings.ToLower("Schema") + "_gen.go"
	defer os.Remove(filePath)
	defer os.Remove("db_gen.go")

	data, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedTextBody, string(data))
}

func TestGenerateDB(t *testing.T) {
	err := generateDB("ayaorm")
	assert.NoError(t, err)

	defer os.Remove("db_gen.go")

	data, err := os.ReadFile("db_gen.go")
	assert.NoError(t, err)
	assert.Equal(t, expectedDBTextBody, string(data))
}

var expectedTextBody = `// Code generated by ayaorm. DO NOT EDIT.
package testss

import (
	"fmt"

	"github.com/okazaki-kk/ayaorm"
)

type CommentRelation struct {
	model *Comment
	*ayaorm.Relation
}

func (m *Comment) newRelation() *CommentRelation {
	r := &CommentRelation{
		m,
		ayaorm.NewRelation(db).SetTable("comments"),
	}
	r.Select(
		"id",
		"content",
		"author",
		"post_id",
		"created_at",
		"updated_at",
	)

	return r
}

func (m Comment) Select(columns ...string) *CommentRelation {
	return m.newRelation().Select(columns...)
}

func (r *CommentRelation) Select(columns ...string) *CommentRelation {
	cs := []string{}
	for _, c := range columns {
		if r.model.isColumnName(c) {
			cs = append(cs, fmt.Sprintf("comments.%s", c))
		} else {
			cs = append(cs, c)
		}
	}
	r.Relation.SetColumns(cs...)
	return r
}

type CommentParams Comment

func (m Comment) Build(p CommentParams) *Comment {
	return &Comment{
		Schema:  ayaorm.Schema{Id: p.Id},
		Content: p.Content,
		Author:  p.Author,
		PostId:  p.PostId,
	}
}

func (u Comment) Create(params CommentParams) (*Comment, error) {
	comment := u.Build(params)
	return u.newRelation().Create(comment)
}

func (r *CommentRelation) Create(comment *Comment) (*Comment, error) {
	err := comment.Save()
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (u *Comment) Update(params CommentParams) error {
	if !ayaorm.IsZero(params.Id) {
		u.Id = params.Id
	}
	if !ayaorm.IsZero(params.Content) {
		u.Content = params.Content
	}
	if !ayaorm.IsZero(params.Author) {
		u.Author = params.Author
	}
	if !ayaorm.IsZero(params.PostId) {
		u.PostId = params.PostId
	}
	return u.Save()
}

func (m *Comment) Save() error {
	lastId, err := m.newRelation().Save()
	if m.Id == 0 {
		m.Id = lastId
	}
	return err
}

func (r *CommentRelation) Save() (int, error) {
	fieldMap := make(map[string]interface{})
	for _, c := range r.Relation.GetColumns() {
		switch c {
		case "content", "comments.content":
			fieldMap["content"] = r.model.Content
		case "author", "comments.author":
			fieldMap["author"] = r.model.Author
		case "post_id", "comments.post_id":
			fieldMap["post_id"] = r.model.PostId
		}
	}

	return r.Relation.Save(r.model.Id, fieldMap)
}

func (m *Comment) Delete() error {
	return m.newRelation().Delete(m.Id)
}

func (m Comment) Count(column ...string) int {
	return m.newRelation().Count(column...)
}

func (m Comment) All() *CommentRelation {
	return m.newRelation()
}

func (m Comment) Limit(limit int) *CommentRelation {
	return m.newRelation().Limit(limit)
}

func (r *CommentRelation) Limit(limit int) *CommentRelation {
	r.Relation.Limit(limit)
	return r
}

func (m Comment) Order(key, order string) *CommentRelation {
	return m.newRelation().Order(key, order)
}

func (r *CommentRelation) Order(key, order string) *CommentRelation {
	r.Relation.Order(key, order)
	return r
}

func (m Comment) Where(column string, conditions ...interface{}) *CommentRelation {
	return m.newRelation().Where(column, conditions...)
}

func (r *CommentRelation) Where(column string, conditions ...interface{}) *CommentRelation {
	r.Relation.Where(column, conditions...)
	return r
}

func (m Comment) Or(column string, conditions ...interface{}) *CommentRelation {
	return m.newRelation().Or(column, conditions...)
}

func (r *CommentRelation) Or(column string, conditions ...interface{}) *CommentRelation {
	r.Relation.Or(column, conditions...)
	return r
}

func (m Comment) GroupBy(columns ...string) *CommentRelation {
	return m.newRelation().GroupBy(columns...)
}

func (r *CommentRelation) GroupBy(columns ...string) *CommentRelation {
	r.Relation.GroupBy(columns...)
	return r
}

func (m Comment) Having(column string, conditions ...interface{}) *CommentRelation {
	return m.newRelation().Having(column, conditions...)
}

func (r *CommentRelation) Having(column string, conditions ...interface{}) *CommentRelation {
	r.Relation.Having(column, conditions...)
	return r
}

func (m Comment) First() (*Comment, error) {
	return m.newRelation().First()
}

func (r *CommentRelation) First() (*Comment, error) {
	r.Relation.First()
	return r.QueryRow()
}

func (m Comment) Last() (*Comment, error) {
	return m.newRelation().Last()
}

func (r *CommentRelation) Last() (*Comment, error) {
	r.Relation.Last()
	return r.QueryRow()
}

func (m Comment) Find(id int) (*Comment, error) {
	return m.newRelation().Find(id)
}

func (r *CommentRelation) Find(id int) (*Comment, error) {
	r.Relation.Find(id)
	return r.QueryRow()
}

func (m Comment) FindBy(column string, value interface{}) (*Comment, error) {
	return m.newRelation().FindBy(column, value)
}

func (r *CommentRelation) FindBy(column string, value interface{}) (*Comment, error) {
	r.Relation.FindBy(column, value)
	return r.QueryRow()
}

func (m Comment) Pluck(column string) ([]interface{}, error) {
	return m.newRelation().Pluck(column)
}

func (r *CommentRelation) Pluck(column string) ([]interface{}, error) {
	return r.Relation.Pluck(column)
}

func (r *CommentRelation) Query() ([]*Comment, error) {
	rows, err := r.Relation.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*Comment{}
	for rows.Next() {
		row := &Comment{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *CommentRelation) QueryRow() (*Comment, error) {
	row := &Comment{}
	err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *Comment) fieldPtrByName(name string) interface{} {
	switch name {
	case "id", "comments.id":
		return &m.Id
	case "content", "comments.content":
		return &m.Content
	case "author", "comments.author":
		return &m.Author
	case "post_id", "comments.post_id":
		return &m.PostId
	case "created_at", "comments.created_at":
		return &m.CreatedAt
	case "updated_at", "comments.updated_at":
		return &m.UpdatedAt
	default:
		return nil
	}
}

func (m *Comment) fieldValuesByName(name string) interface{} {
	switch name {
	case "id", "comments.id":
		return m.Id
	case "content", "comments.content":
		return m.Content
	case "author", "comments.author":
		return m.Author
	case "post_id", "comments.post_id":
		return m.PostId
	case "created_at", "comments.created_at":
		return m.CreatedAt
	case "updated_at", "comments.updated_at":
		return m.UpdatedAt
	default:
		return nil
	}
}

func (m *Comment) fieldPtrsByName(names []string) []interface{} {
	fields := []interface{}{}
	for _, n := range names {
		f := m.fieldPtrByName(n)
		fields = append(fields, f)
	}
	return fields
}

func (m *Comment) isColumnName(name string) bool {
	for _, c := range m.columnNames() {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Comment) columnNames() []string {
	return []string{
		"id",
		"content",
		"author",
		"post_id",
		"created_at",
		"updated_at",
	}
}

func (m Post) Comments() ([]*Comment, error) {
	m.hasManyComments()
	c, err := Comment{}.Where("post_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (m *Post) DeleteDependent() error {
	comments, err := m.Comments()
	if err != nil {
		return err
	}
	for _, comment := range comments {
		err := comment.Delete()
		if err != nil {
			return err
		}
	}
	err = m.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u Post) JoinComments() *PostRelation {
	return u.newRelation().JoinComments()
}

func (u *PostRelation) JoinComments() *PostRelation {
	u.Relation.InnerJoin("posts", "comments", true)
	return u
}

func (u Comment) Post() (*Post, error) {
	u.belongsToPost()
	return Post{}.Find(u.PostId)
}

func (u Comment) JoinPost() *CommentRelation {
	return u.newRelation().JoinPost()
}

func (u *CommentRelation) JoinPost() *CommentRelation {
	u.Relation.InnerJoin("comments", "posts", false)
	return u
}

func (m Post) IsValid() (bool, []error) {
	result := true
	var errors []error

	rules := map[string]*ayaorm.Validation{
		"author":  m.validatesPresenceOfAuthor().Rule(),
		"content": m.validateLengthOfContent().Rule(),
	}

	for name, rule := range rules {
		if ok, errs := ayaorm.NewValidator(rule).IsValid(name, m.fieldValuesByName(name)); !ok {
			result = false
			errors = append(errors, errs...)
		}
	}

	if len(errors) > 0 {
		result = false
	}
	return result, errors
}

func (m User) IsValid() (bool, []error) {
	result := true
	var errors []error

	rules := map[string]*ayaorm.Validation{
		"age": m.validateNumericalityOfAge().Rule(),
	}

	for name, rule := range rules {
		if ok, errs := ayaorm.NewValidator(rule).IsValid(name, m.fieldValuesByName(name)); !ok {
			result = false
			errors = append(errors, errs...)
		}
	}

	if len(errors) > 0 {
		result = false
	}
	return result, errors
}
`

var expectedDBTextBody = `// Code generated by ayaorm. DO NOT EDIT.
package ayaorm

import "database/sql"

var db *sql.DB
`

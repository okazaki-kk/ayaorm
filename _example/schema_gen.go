// Code generated by ayaorm. DO NOT EDIT.
package main

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
	return u.newRelation().Update(u.Id, params)
}

func (r *CommentRelation) Update(id int, params CommentParams) error {
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
	return r.Relation.Update(id, fieldMap)
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

	return r.Relation.Save(fieldMap)
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

type PostRelation struct {
	model *Post
	*ayaorm.Relation
}

func (m *Post) newRelation() *PostRelation {
	r := &PostRelation{
		m,
		ayaorm.NewRelation(db).SetTable("posts"),
	}
	r.Select(
		"id",
		"content",
		"author",
		"created_at",
		"updated_at",
	)

	return r
}

func (m Post) Select(columns ...string) *PostRelation {
	return m.newRelation().Select(columns...)
}

func (r *PostRelation) Select(columns ...string) *PostRelation {
	cs := []string{}
	for _, c := range columns {
		if r.model.isColumnName(c) {
			cs = append(cs, fmt.Sprintf("posts.%s", c))
		} else {
			cs = append(cs, c)
		}
	}
	r.Relation.SetColumns(cs...)
	return r
}

type PostParams Post

func (m Post) Build(p PostParams) *Post {
	return &Post{
		Schema:  ayaorm.Schema{Id: p.Id},
		Content: p.Content,
		Author:  p.Author,
	}
}

func (u Post) Create(params PostParams) (*Post, error) {
	post := u.Build(params)
	return u.newRelation().Create(post)
}

func (r *PostRelation) Create(post *Post) (*Post, error) {
	err := post.Save()
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *Post) Update(params PostParams) error {
	return u.newRelation().Update(u.Id, params)
}

func (r *PostRelation) Update(id int, params PostParams) error {
	fieldMap := make(map[string]interface{})
	for _, c := range r.Relation.GetColumns() {
		switch c {
		case "content", "posts.content":
			fieldMap["content"] = r.model.Content
		case "author", "posts.author":
			fieldMap["author"] = r.model.Author
		}
	}
	return r.Relation.Update(id, fieldMap)
}

func (m *Post) Save() error {
	lastId, err := m.newRelation().Save()
	if m.Id == 0 {
		m.Id = lastId
	}
	return err
}

func (r *PostRelation) Save() (int, error) {
	fieldMap := make(map[string]interface{})
	for _, c := range r.Relation.GetColumns() {
		switch c {
		case "content", "posts.content":
			fieldMap["content"] = r.model.Content
		case "author", "posts.author":
			fieldMap["author"] = r.model.Author
		}
	}

	return r.Relation.Save(fieldMap)
}

func (m *Post) Delete() error {
	return m.newRelation().Delete(m.Id)
}

func (m Post) Count(column ...string) int {
	return m.newRelation().Count(column...)
}

func (m Post) All() *PostRelation {
	return m.newRelation()
}

func (m Post) Limit(limit int) *PostRelation {
	return m.newRelation().Limit(limit)
}

func (r *PostRelation) Limit(limit int) *PostRelation {
	r.Relation.Limit(limit)
	return r
}

func (m Post) Order(key, order string) *PostRelation {
	return m.newRelation().Order(key, order)
}

func (r *PostRelation) Order(key, order string) *PostRelation {
	r.Relation.Order(key, order)
	return r
}

func (m Post) Where(column string, conditions ...interface{}) *PostRelation {
	return m.newRelation().Where(column, conditions...)
}

func (r *PostRelation) Where(column string, conditions ...interface{}) *PostRelation {
	r.Relation.Where(column, conditions...)
	return r
}

func (m Post) First() (*Post, error) {
	return m.newRelation().First()
}

func (r *PostRelation) First() (*Post, error) {
	r.Relation.First()
	return r.QueryRow()
}

func (m Post) Last() (*Post, error) {
	return m.newRelation().Last()
}

func (r *PostRelation) Last() (*Post, error) {
	r.Relation.Last()
	return r.QueryRow()
}

func (m Post) Find(id int) (*Post, error) {
	return m.newRelation().Find(id)
}

func (r *PostRelation) Find(id int) (*Post, error) {
	r.Relation.Find(id)
	return r.QueryRow()
}

func (m Post) FindBy(column string, value interface{}) (*Post, error) {
	return m.newRelation().FindBy(column, value)
}

func (r *PostRelation) FindBy(column string, value interface{}) (*Post, error) {
	r.Relation.FindBy(column, value)
	return r.QueryRow()
}

func (m Post) Pluck(column string) ([]interface{}, error) {
	return m.newRelation().Pluck(column)
}

func (r *PostRelation) Pluck(column string) ([]interface{}, error) {
	return r.Relation.Pluck(column)
}

func (r *PostRelation) Query() ([]*Post, error) {
	rows, err := r.Relation.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*Post{}
	for rows.Next() {
		row := &Post{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *PostRelation) QueryRow() (*Post, error) {
	row := &Post{}
	err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *Post) fieldPtrByName(name string) interface{} {
	switch name {
	case "id", "posts.id":
		return &m.Id
	case "content", "posts.content":
		return &m.Content
	case "author", "posts.author":
		return &m.Author
	case "created_at", "posts.created_at":
		return &m.CreatedAt
	case "updated_at", "posts.updated_at":
		return &m.UpdatedAt
	default:
		return nil
	}
}

func (m *Post) fieldPtrsByName(names []string) []interface{} {
	fields := []interface{}{}
	for _, n := range names {
		f := m.fieldPtrByName(n)
		fields = append(fields, f)
	}
	return fields
}

func (m *Post) isColumnName(name string) bool {
	for _, c := range m.columnNames() {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Post) columnNames() []string {
	return []string{
		"id",
		"content",
		"author",
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
	u.Relation.InnerJoin("posts", "comments", true)
	return u
}

func (u Comment) Post() (*Post, error) {
	return Post{}.Find(u.PostId)
}

func (u Comment) JoinPost() *CommentRelation {
	return u.newRelation().JoinPost()
}

func (u *CommentRelation) JoinPost() *CommentRelation {
	u.Relation.InnerJoin("comments", "posts", false)
	return u
}

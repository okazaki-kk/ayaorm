// Code generated by ayaorm. DO NOT EDIT.
package test_sqlite

import (
	"errors"
	"fmt"

	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/utils"
	"github.com/okazaki-kk/ayaorm/validate"
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
		"achievement_rate",
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
		Schema:          ayaorm.Schema{Id: p.Id},
		Content:         p.Content,
		Author:          p.Author,
		AchievementRate: p.AchievementRate,
		PostId:          p.PostId,
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

func (u Comment) CreateAll(params []CommentParams) error {
	comments := make([]*Comment, len(params))
	for i, p := range params {
		comments[i] = u.Build(p)
	}
	return u.newRelation().CreateAll(comments)
}

func (r *CommentRelation) CreateAll(comments []*Comment) error {
	fieldMap := make(map[string][]interface{})
	for _, comment := range comments {
		for _, c := range r.Relation.GetColumns() {
			switch c {
			case "content", "comments.content":
				fieldMap["content"] = append(fieldMap["content"], comment.Content)
			case "author", "comments.author":
				fieldMap["author"] = append(fieldMap["author"], comment.Author)
			case "achievement_rate", "comments.achievement_rate":
				fieldMap["achievement_rate"] = append(fieldMap["achievement_rate"], comment.AchievementRate)
			case "post_id", "comments.post_id":
				fieldMap["post_id"] = append(fieldMap["post_id"], comment.PostId)
			}
		}
	}
	return r.Relation.CreateAll(fieldMap)
}

func (u *Comment) Update(params CommentParams) error {
	if !utils.IsZero(params.Id) {
		u.Id = params.Id
	}
	if !utils.IsZero(params.Content) {
		u.Content = params.Content
	}
	if !utils.IsZero(params.Author) {
		u.Author = params.Author
	}
	if !utils.IsZero(params.AchievementRate) {
		u.AchievementRate = params.AchievementRate
	}
	if !utils.IsZero(params.PostId) {
		u.PostId = params.PostId
	}
	return u.Save()
}

func (m *Comment) Save() error {
	ok, errs := m.IsValid()
	if !ok {
		return errors.Join(errs...)
	}

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
		case "achievement_rate", "comments.achievement_rate":
			fieldMap["achievement_rate"] = r.model.AchievementRate
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
	case "achievement_rate", "comments.achievement_rate":
		return &m.AchievementRate
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
	case "achievement_rate", "comments.achievement_rate":
		return m.AchievementRate
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
		"achievement_rate",
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

func (u Post) CreateAll(params []PostParams) error {
	posts := make([]*Post, len(params))
	for i, p := range params {
		posts[i] = u.Build(p)
	}
	return u.newRelation().CreateAll(posts)
}

func (r *PostRelation) CreateAll(posts []*Post) error {
	fieldMap := make(map[string][]interface{})
	for _, post := range posts {
		for _, c := range r.Relation.GetColumns() {
			switch c {
			case "content", "posts.content":
				fieldMap["content"] = append(fieldMap["content"], post.Content)
			case "author", "posts.author":
				fieldMap["author"] = append(fieldMap["author"], post.Author)
			}
		}
	}
	return r.Relation.CreateAll(fieldMap)
}

func (u *Post) Update(params PostParams) error {
	if !utils.IsZero(params.Id) {
		u.Id = params.Id
	}
	if !utils.IsZero(params.Content) {
		u.Content = params.Content
	}
	if !utils.IsZero(params.Author) {
		u.Author = params.Author
	}
	return u.Save()
}

func (m *Post) Save() error {
	ok, errs := m.IsValid()
	if !ok {
		return errors.Join(errs...)
	}

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

	return r.Relation.Save(r.model.Id, fieldMap)
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

func (m Post) Or(column string, conditions ...interface{}) *PostRelation {
	return m.newRelation().Or(column, conditions...)
}

func (r *PostRelation) Or(column string, conditions ...interface{}) *PostRelation {
	r.Relation.Or(column, conditions...)
	return r
}

func (m Post) GroupBy(columns ...string) *PostRelation {
	return m.newRelation().GroupBy(columns...)
}

func (r *PostRelation) GroupBy(columns ...string) *PostRelation {
	r.Relation.GroupBy(columns...)
	return r
}

func (m Post) Having(column string, conditions ...interface{}) *PostRelation {
	return m.newRelation().Having(column, conditions...)
}

func (r *PostRelation) Having(column string, conditions ...interface{}) *PostRelation {
	r.Relation.Having(column, conditions...)
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

func (m *Post) fieldValuesByName(name string) interface{} {
	switch name {
	case "id", "posts.id":
		return m.Id
	case "content", "posts.content":
		return m.Content
	case "author", "posts.author":
		return m.Author
	case "created_at", "posts.created_at":
		return m.CreatedAt
	case "updated_at", "posts.updated_at":
		return m.UpdatedAt
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

type ProjectRelation struct {
	model *Project
	*ayaorm.Relation
}

func (m *Project) newRelation() *ProjectRelation {
	r := &ProjectRelation{
		m,
		ayaorm.NewRelation(db).SetTable("projects"),
	}
	r.Select(
		"id",
		"name",
		"post_id",
		"created_at",
		"updated_at",
	)

	return r
}

func (m Project) Select(columns ...string) *ProjectRelation {
	return m.newRelation().Select(columns...)
}

func (r *ProjectRelation) Select(columns ...string) *ProjectRelation {
	cs := []string{}
	for _, c := range columns {
		if r.model.isColumnName(c) {
			cs = append(cs, fmt.Sprintf("projects.%s", c))
		} else {
			cs = append(cs, c)
		}
	}
	r.Relation.SetColumns(cs...)
	return r
}

type ProjectParams Project

func (m Project) Build(p ProjectParams) *Project {
	return &Project{
		Schema: ayaorm.Schema{Id: p.Id},
		Name:   p.Name,
		PostId: p.PostId,
	}
}

func (u Project) Create(params ProjectParams) (*Project, error) {
	project := u.Build(params)
	return u.newRelation().Create(project)
}

func (r *ProjectRelation) Create(project *Project) (*Project, error) {
	err := project.Save()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (u Project) CreateAll(params []ProjectParams) error {
	projects := make([]*Project, len(params))
	for i, p := range params {
		projects[i] = u.Build(p)
	}
	return u.newRelation().CreateAll(projects)
}

func (r *ProjectRelation) CreateAll(projects []*Project) error {
	fieldMap := make(map[string][]interface{})
	for _, project := range projects {
		for _, c := range r.Relation.GetColumns() {
			switch c {
			case "name", "projects.name":
				fieldMap["name"] = append(fieldMap["name"], project.Name)
			case "post_id", "projects.post_id":
				fieldMap["post_id"] = append(fieldMap["post_id"], project.PostId)
			}
		}
	}
	return r.Relation.CreateAll(fieldMap)
}

func (u *Project) Update(params ProjectParams) error {
	if !utils.IsZero(params.Id) {
		u.Id = params.Id
	}
	if !utils.IsZero(params.Name) {
		u.Name = params.Name
	}
	if !utils.IsZero(params.PostId) {
		u.PostId = params.PostId
	}
	return u.Save()
}

func (m *Project) Save() error {
	ok, errs := m.IsValid()
	if !ok {
		return errors.Join(errs...)
	}

	lastId, err := m.newRelation().Save()
	if m.Id == 0 {
		m.Id = lastId
	}
	return err
}

func (r *ProjectRelation) Save() (int, error) {
	fieldMap := make(map[string]interface{})
	for _, c := range r.Relation.GetColumns() {
		switch c {
		case "name", "projects.name":
			fieldMap["name"] = r.model.Name
		case "post_id", "projects.post_id":
			fieldMap["post_id"] = r.model.PostId
		}
	}

	return r.Relation.Save(r.model.Id, fieldMap)
}

func (m *Project) Delete() error {
	return m.newRelation().Delete(m.Id)
}

func (m Project) Count(column ...string) int {
	return m.newRelation().Count(column...)
}

func (m Project) All() *ProjectRelation {
	return m.newRelation()
}

func (m Project) Limit(limit int) *ProjectRelation {
	return m.newRelation().Limit(limit)
}

func (r *ProjectRelation) Limit(limit int) *ProjectRelation {
	r.Relation.Limit(limit)
	return r
}

func (m Project) Order(key, order string) *ProjectRelation {
	return m.newRelation().Order(key, order)
}

func (r *ProjectRelation) Order(key, order string) *ProjectRelation {
	r.Relation.Order(key, order)
	return r
}

func (m Project) Where(column string, conditions ...interface{}) *ProjectRelation {
	return m.newRelation().Where(column, conditions...)
}

func (r *ProjectRelation) Where(column string, conditions ...interface{}) *ProjectRelation {
	r.Relation.Where(column, conditions...)
	return r
}

func (m Project) Or(column string, conditions ...interface{}) *ProjectRelation {
	return m.newRelation().Or(column, conditions...)
}

func (r *ProjectRelation) Or(column string, conditions ...interface{}) *ProjectRelation {
	r.Relation.Or(column, conditions...)
	return r
}

func (m Project) GroupBy(columns ...string) *ProjectRelation {
	return m.newRelation().GroupBy(columns...)
}

func (r *ProjectRelation) GroupBy(columns ...string) *ProjectRelation {
	r.Relation.GroupBy(columns...)
	return r
}

func (m Project) Having(column string, conditions ...interface{}) *ProjectRelation {
	return m.newRelation().Having(column, conditions...)
}

func (r *ProjectRelation) Having(column string, conditions ...interface{}) *ProjectRelation {
	r.Relation.Having(column, conditions...)
	return r
}

func (m Project) First() (*Project, error) {
	return m.newRelation().First()
}

func (r *ProjectRelation) First() (*Project, error) {
	r.Relation.First()
	return r.QueryRow()
}

func (m Project) Last() (*Project, error) {
	return m.newRelation().Last()
}

func (r *ProjectRelation) Last() (*Project, error) {
	r.Relation.Last()
	return r.QueryRow()
}

func (m Project) Find(id int) (*Project, error) {
	return m.newRelation().Find(id)
}

func (r *ProjectRelation) Find(id int) (*Project, error) {
	r.Relation.Find(id)
	return r.QueryRow()
}

func (m Project) FindBy(column string, value interface{}) (*Project, error) {
	return m.newRelation().FindBy(column, value)
}

func (r *ProjectRelation) FindBy(column string, value interface{}) (*Project, error) {
	r.Relation.FindBy(column, value)
	return r.QueryRow()
}

func (m Project) Pluck(column string) ([]interface{}, error) {
	return m.newRelation().Pluck(column)
}

func (r *ProjectRelation) Pluck(column string) ([]interface{}, error) {
	return r.Relation.Pluck(column)
}

func (r *ProjectRelation) Query() ([]*Project, error) {
	rows, err := r.Relation.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*Project{}
	for rows.Next() {
		row := &Project{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *ProjectRelation) QueryRow() (*Project, error) {
	row := &Project{}
	err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m *Project) fieldPtrByName(name string) interface{} {
	switch name {
	case "id", "projects.id":
		return &m.Id
	case "name", "projects.name":
		return &m.Name
	case "post_id", "projects.post_id":
		return &m.PostId
	case "created_at", "projects.created_at":
		return &m.CreatedAt
	case "updated_at", "projects.updated_at":
		return &m.UpdatedAt
	default:
		return nil
	}
}

func (m *Project) fieldValuesByName(name string) interface{} {
	switch name {
	case "id", "projects.id":
		return m.Id
	case "name", "projects.name":
		return m.Name
	case "post_id", "projects.post_id":
		return m.PostId
	case "created_at", "projects.created_at":
		return m.CreatedAt
	case "updated_at", "projects.updated_at":
		return m.UpdatedAt
	default:
		return nil
	}
}

func (m *Project) fieldPtrsByName(names []string) []interface{} {
	fields := []interface{}{}
	for _, n := range names {
		f := m.fieldPtrByName(n)
		fields = append(fields, f)
	}
	return fields
}

func (m *Project) isColumnName(name string) bool {
	for _, c := range m.columnNames() {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Project) columnNames() []string {
	return []string{
		"id",
		"name",
		"post_id",
		"created_at",
		"updated_at",
	}
}

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
		"age1",
		"age2",
		"address",
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
		Schema:  ayaorm.Schema{Id: p.Id},
		Name:    p.Name,
		Age:     p.Age,
		Age1:    p.Age1,
		Age2:    p.Age2,
		Address: p.Address,
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

func (u User) CreateAll(params []UserParams) error {
	users := make([]*User, len(params))
	for i, p := range params {
		users[i] = u.Build(p)
	}
	return u.newRelation().CreateAll(users)
}

func (r *UserRelation) CreateAll(users []*User) error {
	fieldMap := make(map[string][]interface{})
	for _, user := range users {
		for _, c := range r.Relation.GetColumns() {
			switch c {
			case "name", "users.name":
				fieldMap["name"] = append(fieldMap["name"], user.Name)
			case "age", "users.age":
				fieldMap["age"] = append(fieldMap["age"], user.Age)
			case "age1", "users.age1":
				fieldMap["age1"] = append(fieldMap["age1"], user.Age1)
			case "age2", "users.age2":
				fieldMap["age2"] = append(fieldMap["age2"], user.Age2)
			case "address", "users.address":
				fieldMap["address"] = append(fieldMap["address"], user.Address)
			}
		}
	}
	return r.Relation.CreateAll(fieldMap)
}

func (u *User) Update(params UserParams) error {
	if !utils.IsZero(params.Id) {
		u.Id = params.Id
	}
	if !utils.IsZero(params.Name) {
		u.Name = params.Name
	}
	if !utils.IsZero(params.Age) {
		u.Age = params.Age
	}
	if !utils.IsZero(params.Age1) {
		u.Age1 = params.Age1
	}
	if !utils.IsZero(params.Age2) {
		u.Age2 = params.Age2
	}
	if !utils.IsZero(params.Address) {
		u.Address = params.Address
	}
	return u.Save()
}

func (m *User) Save() error {
	ok, errs := m.IsValid()
	if !ok {
		return errors.Join(errs...)
	}

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
		case "age1", "users.age1":
			fieldMap["age1"] = r.model.Age1
		case "age2", "users.age2":
			fieldMap["age2"] = r.model.Age2
		case "address", "users.address":
			fieldMap["address"] = r.model.Address
		}
	}

	return r.Relation.Save(r.model.Id, fieldMap)
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

func (m User) Where(column string, conditions ...interface{}) *UserRelation {
	return m.newRelation().Where(column, conditions...)
}

func (r *UserRelation) Where(column string, conditions ...interface{}) *UserRelation {
	r.Relation.Where(column, conditions...)
	return r
}

func (m User) Or(column string, conditions ...interface{}) *UserRelation {
	return m.newRelation().Or(column, conditions...)
}

func (r *UserRelation) Or(column string, conditions ...interface{}) *UserRelation {
	r.Relation.Or(column, conditions...)
	return r
}

func (m User) GroupBy(columns ...string) *UserRelation {
	return m.newRelation().GroupBy(columns...)
}

func (r *UserRelation) GroupBy(columns ...string) *UserRelation {
	r.Relation.GroupBy(columns...)
	return r
}

func (m User) Having(column string, conditions ...interface{}) *UserRelation {
	return m.newRelation().Having(column, conditions...)
}

func (r *UserRelation) Having(column string, conditions ...interface{}) *UserRelation {
	r.Relation.Having(column, conditions...)
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
	case "age1", "users.age1":
		return &m.Age1
	case "age2", "users.age2":
		return &m.Age2
	case "address", "users.address":
		return &m.Address
	case "created_at", "users.created_at":
		return &m.CreatedAt
	case "updated_at", "users.updated_at":
		return &m.UpdatedAt
	default:
		return nil
	}
}

func (m *User) fieldValuesByName(name string) interface{} {
	switch name {
	case "id", "users.id":
		return m.Id
	case "name", "users.name":
		return m.Name
	case "age", "users.age":
		return m.Age
	case "age1", "users.age1":
		return m.Age1
	case "age2", "users.age2":
		return m.Age2
	case "address", "users.address":
		return m.Address
	case "created_at", "users.created_at":
		return m.CreatedAt
	case "updated_at", "users.updated_at":
		return m.UpdatedAt
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
		"age1",
		"age2",
		"address",
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

func (u Project) Post() (*Post, error) {
	u.belongsToPost()
	return Post{}.Find(u.PostId)
}

func (u Project) JoinPost() *ProjectRelation {
	return u.newRelation().JoinPost()
}

func (u *ProjectRelation) JoinPost() *ProjectRelation {
	u.Relation.InnerJoin("projects", "posts", false)
	return u
}

func (u Post) Project() (*Project, error) {
	u.hasOneProject()
	c, err := Project{}.FindBy("post_id", u.Id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (m Comment) IsValid() (bool, []error) {
	result := true
	var errors []error

	var on validate.On
	if utils.IsZero(m.Id) {
		on = validate.On{OnCreate: true, OnUpdate: false}
	} else {
		on = validate.On{OnCreate: false, OnUpdate: true}
	}

	rules := map[string]*validate.Validation{}

	for name, rule := range rules {
		if ok, errs := validate.NewValidator(rule).On(on).IsValid(name, m.fieldValuesByName(name)); !ok {
			result = false
			errors = append(errors, errs...)
		}
	}

	if len(errors) > 0 {
		result = false
	}
	return result, errors
}

func (m Post) IsValid() (bool, []error) {
	result := true
	var errors []error

	var on validate.On
	if utils.IsZero(m.Id) {
		on = validate.On{OnCreate: true, OnUpdate: false}
	} else {
		on = validate.On{OnCreate: false, OnUpdate: true}
	}

	rules := map[string]*validate.Validation{
		"author":  m.validatesPresenceOfAuthor().Rule(),
		"content": m.validateLengthOfContent().Rule(),
	}

	for name, rule := range rules {
		if ok, errs := validate.NewValidator(rule).On(on).IsValid(name, m.fieldValuesByName(name)); !ok {
			result = false
			errors = append(errors, errs...)
		}
	}

	if len(errors) > 0 {
		result = false
	}
	return result, errors
}

func (m Project) IsValid() (bool, []error) {
	result := true
	var errors []error

	var on validate.On
	if utils.IsZero(m.Id) {
		on = validate.On{OnCreate: true, OnUpdate: false}
	} else {
		on = validate.On{OnCreate: false, OnUpdate: true}
	}

	rules := map[string]*validate.Validation{}

	for name, rule := range rules {
		if ok, errs := validate.NewValidator(rule).On(on).IsValid(name, m.fieldValuesByName(name)); !ok {
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

	var on validate.On
	if utils.IsZero(m.Id) {
		on = validate.On{OnCreate: true, OnUpdate: false}
	} else {
		on = validate.On{OnCreate: false, OnUpdate: true}
	}

	rules := map[string]*validate.Validation{
		"age":  m.validateNumericalityOfAge().Rule(),
		"age1": m.validateNumericalityOfAge1().Rule(),
		"age2": m.validateNumericalityOfAge2().Rule(),
	}

	for name, rule := range rules {
		if ok, errs := validate.NewValidator(rule).On(on).IsValid(name, m.fieldValuesByName(name)); !ok {
			result = false
			errors = append(errors, errs...)
		}
	}

	customs := []*validate.Validation{m.validateCustomRule().Rule()}
	for _, rule := range customs {
		custom := validate.NewValidator(rule).Custom()
		custom(&errors)
	}

	if len(errors) > 0 {
		result = false
	}
	return result, errors
}

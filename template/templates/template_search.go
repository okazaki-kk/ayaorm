package templates

var SearchTextBody = `
	{{define "Search"}}
	func (m {{.ModelName}}) Count(column ...string) int {
		return m.newRelation().Count(column...)
	}

	func (m {{.ModelName}}) All() *{{.ModelName}}Relation {
		return m.newRelation()
	}

	func (m {{.ModelName}}) Limit(limit int) *{{.ModelName}}Relation {
		return m.newRelation().Limit(limit)
	}

	func (r *{{.ModelName}}Relation) Limit(limit int) *{{.ModelName}}Relation {
		r.Relation.Limit(limit)
		return r
	}

	func (m {{.ModelName}}) Order(key, order string) *{{.ModelName}}Relation {
		return m.newRelation().Order(key, order)
	}

	func (r *{{.ModelName}}Relation) Order(key, order string) *{{.ModelName}}Relation {
		r.Relation.Order(key, order)
		return r
	}

	func (m {{.ModelName}}) Where(column string, conditions ...interface{}) *{{.ModelName}}Relation {
		return m.newRelation().Where(column, conditions...)
	}

	func (r *{{.ModelName}}Relation) Where(column string, conditions ...interface{}) *{{.ModelName}}Relation {
		r.Relation.Where(column, conditions...)
		return r
	}

	func (m {{.ModelName}}) Or(column string, conditions ...interface{}) *{{.ModelName}}Relation {
		return m.newRelation().Or(column, conditions...)
	}

	func (r *{{.ModelName}}Relation) Or(column string, conditions ...interface{}) *{{.ModelName}}Relation {
		r.Relation.Or(column, conditions...)
		return r
	}

	func (m {{.ModelName}}) GroupBy(columns ...string) *{{.ModelName}}Relation {
		return m.newRelation().GroupBy(columns...)
	}

	func (r *{{.ModelName}}Relation) GroupBy(columns ...string) *{{.ModelName}}Relation {
		r.Relation.GroupBy(columns...)
		return r
	}

	func (m {{.ModelName}}) First() (*{{.ModelName}}, error) {
		return m.newRelation().First()
	}

	func (r *{{.ModelName}}Relation) First() (*{{.ModelName}}, error) {
		r.Relation.First()
		return r.QueryRow()
	}

	func (m {{.ModelName}}) Last() (*{{.ModelName}}, error) {
		return m.newRelation().Last()
	}

	func (r *{{.ModelName}}Relation) Last() (*{{.ModelName}}, error) {
		r.Relation.Last()
		return r.QueryRow()
	}

	func (m {{.ModelName}}) Find(id int) (*{{.ModelName}}, error) {
		return m.newRelation().Find(id)
	}

	func (r *{{.ModelName}}Relation) Find(id int) (*{{.ModelName}}, error) {
		r.Relation.Find(id)
		return r.QueryRow()
	}

	func (m {{.ModelName}}) FindBy(column string, value interface{}) (*{{.ModelName}}, error) {
		return m.newRelation().FindBy(column, value)
	}

	func (r *{{.ModelName}}Relation) FindBy(column string, value interface{}) (*{{.ModelName}}, error) {
		r.Relation.FindBy(column, value)
		return r.QueryRow()
	}

	func (m {{.ModelName}}) Pluck(column string) ([]interface{}, error) {
		return m.newRelation().Pluck(column)
	}

	func (r *{{.ModelName}}Relation) Pluck(column string) ([]interface{}, error) {
		return r.Relation.Pluck(column)
	}
	{{end}}
`

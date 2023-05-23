package template

var searchTextBody = `
	{{define "Search"}}
	func (m {{.modelName}}) Count(column ...string) int {
		return m.newRelation().Count(column...)
	}

	func (m {{.modelName}}) All() *{{.modelName}}Relation {
		return m.newRelation()
	}

	func (m {{.modelName}}) Limit(limit int) *{{.modelName}}Relation {
		return m.newRelation().Limit(limit)
	}

	func (r *{{.modelName}}Relation) Limit(limit int) *{{.modelName}}Relation {
		r.Relation.Limit(limit)
		return r
	}

	func (m {{.modelName}}) Order(key, order string) *{{.modelName}}Relation {
		return m.newRelation().Order(key, order)
	}

	func (r *{{.modelName}}Relation) Order(key, order string) *{{.modelName}}Relation {
		r.Relation.Order(key, order)
		return r
	}

	func (m {{.modelName}}) Where(column string, value interface{}) *{{.modelName}}Relation {
		return m.newRelation().Where(column, value)
	}

	func (r *{{.modelName}}Relation) Where(column string, value interface{}) *{{.modelName}}Relation {
		r.Relation.Where(column, value)
		return r
	}

	func (m {{.modelName}}) First() (*{{.modelName}}, error) {
		return m.newRelation().First()
	}

	func (r *{{.modelName}}Relation) First() (*{{.modelName}}, error) {
		r.Relation.First()
		return r.QueryRow()
	}

	func (m {{.modelName}}) Last() (*{{.modelName}}, error) {
		return m.newRelation().Last()
	}

	func (r *{{.modelName}}Relation) Last() (*{{.modelName}}, error) {
		r.Relation.Last()
		return r.QueryRow()
	}

	func (m {{.modelName}}) Find(id int) (*{{.modelName}}, error) {
		return m.newRelation().Find(id)
	}

	func (r *{{.modelName}}Relation) Find(id int) (*{{.modelName}}, error) {
		r.Relation.Find(id)
		return r.QueryRow()
	}

	func (m {{.modelName}}) FindBy(column string, value interface{}) (*{{.modelName}}, error) {
		return m.newRelation().FindBy(column, value)
	}

	func (r *{{.modelName}}Relation) FindBy(column string, value interface{}) (*{{.modelName}}, error) {
		r.Relation.FindBy(column, value)
		return r.QueryRow()
	}

	func (m {{.modelName}}) Pluck(column string) ([]interface{}, error) {
		return m.newRelation().Pluck(column)
	}

	func (r *{{.modelName}}Relation) Pluck(column string) ([]interface{}, error) {
		return r.Relation.Pluck(column)
	}
	{{end}}
`

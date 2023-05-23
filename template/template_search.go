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

	{{end}}
`

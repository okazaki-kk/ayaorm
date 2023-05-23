package templates

var CrudTextBody = `
	{{define "CRUD"}}
	func (m {{.modelName}}) Build(p {{.modelName}}Params) *{{.modelName}} {
		return &{{.modelName}}{
			Schema: ayaorm.Schema{Id: p.Id},
			{{ range $column := .columns -}}
			{{ if eq $column "CreatedAt" -}}
			{{ continue }}
			{{ end -}}
			{{ if eq $column "UpdatedAt" -}}
			{{ continue }}
			{{ end -}}
			{{ if eq $column "Id" -}}
			{{ continue }}
			{{ end -}}
			{{ $column }}: p.{{ $column }},
			{{ end -}}
		}
	}

	func (u {{.modelName}}) Create(params {{.modelName}}Params) (*{{.modelName}}, error) {
		{{toSnakeCase .modelName}} := u.Build(params)
		return u.newRelation().Create({{toSnakeCase .modelName}})
	}

	func (r *{{.modelName}}Relation) Create({{toSnakeCase .modelName}} *{{.modelName}}) (*{{.modelName}}, error) {
		err := {{toSnakeCase .modelName}}.Save()
		if err != nil {
			return nil, err
		}
		return {{toSnakeCase .modelName}}, nil
	}

	func (u *{{.modelName}}) Update(params {{.modelName}}Params) error {
		return u.newRelation().Update(u.Id, params)
	}

	func (r *{{.modelName}}Relation) Update(id int, params {{.modelName}}Params) error {
		fieldMap := make(map[string]interface{})
		for _, c := range r.Relation.GetColumns() {
			switch c {
				{{ range $column := .columns -}}
				{{ if eq $column "Id" -}}
				{{ continue }}
				{{ end -}}
				{{ if eq $column "CreatedAt" -}}
				{{ continue }}
				{{ end -}}
				{{ if eq $column "UpdatedAt" -}}
				{{ continue }}
				{{ end -}}
				case "{{ toSnakeCase  $column}}", "{{$.snakeCaseModelName}}.{{toSnakeCase $column}}":
					fieldMap["{{toSnakeCase $column}}"] = r.model.{{$column}}
				{{ end -}}
			}
		}
		return r.Relation.Update(id, fieldMap)
	}

	func (m *{{.modelName}}) Save() error {
		lastId, err := m.newRelation().Save()
		if m.Id == 0 {
			m.Id = lastId
		}
		return err
	}

	func (r *{{.modelName}}Relation) Save() (int, error) {
		fieldMap := make(map[string]interface{})
		for _, c := range r.Relation.GetColumns() {
			switch c {
				{{ range $column := .columns -}}
				{{ if eq $column "Id" -}}
				{{ continue }}
				{{ end -}}
				{{ if eq $column "CreatedAt" -}}
				{{ continue }}
				{{ end -}}
				{{ if eq $column "UpdatedAt" -}}
				{{ continue }}
				{{ end -}}
				case "{{ toSnakeCase  $column}}", "{{$.snakeCaseModelName}}.{{toSnakeCase $column}}":
					fieldMap["{{toSnakeCase $column}}"] = r.model.{{$column}}
				{{ end -}}
			}
		}

		return r.Relation.Save(fieldMap)
	}

	func (m *{{.modelName}}) Delete() error {
		return m.newRelation().Delete(m.Id)
	}
	{{end}}
`

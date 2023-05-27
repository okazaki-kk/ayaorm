package templates

var CrudTextBody = `
	{{define "CRUD"}}
	func (m {{.ModelName}}) Build(p {{.ModelName}}Params) *{{.ModelName}} {
		return &{{.ModelName}}{
			Schema: ayaorm.Schema{Id: p.Id},
			{{ range $column := .Columns -}}
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

	func (u {{.ModelName}}) Create(params {{.ModelName}}Params) (*{{.ModelName}}, error) {
		{{toSnakeCase .ModelName}} := u.Build(params)
		return u.newRelation().Create({{toSnakeCase .ModelName}})
	}

	func (r *{{.ModelName}}Relation) Create({{toSnakeCase .ModelName}} *{{.ModelName}}) (*{{.ModelName}}, error) {
		err := {{toSnakeCase .ModelName}}.Save()
		if err != nil {
			return nil, err
		}
		return {{toSnakeCase .ModelName}}, nil
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

	func (m *{{.ModelName}}) Save() error {
		lastId, err := m.newRelation().Save()
		if m.Id == 0 {
			m.Id = lastId
		}
		return err
	}

	func (r *{{.ModelName}}Relation) Save() (int, error) {
		fieldMap := make(map[string]interface{})
		for _, c := range r.Relation.GetColumns() {
			switch c {
				{{ range $column := .Columns -}}
				{{ if eq $column "Id" -}}
				{{ continue }}
				{{ end -}}
				{{ if eq $column "CreatedAt" -}}
				{{ continue }}
				{{ end -}}
				{{ if eq $column "UpdatedAt" -}}
				{{ continue }}
				{{ end -}}
				case "{{ toSnakeCase  $column}}", "{{$.SnakeCaseModelName}}.{{toSnakeCase $column}}":
					fieldMap["{{toSnakeCase $column}}"] = r.model.{{$column}}
				{{ end -}}
			}
		}

		return r.Relation.Save(r.model.Id, fieldMap)
	}

	func (m *{{.ModelName}}) Delete() error {
		return m.newRelation().Delete(m.Id)
	}
	{{end}}
`

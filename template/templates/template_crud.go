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

	func (u {{.ModelName}}) CreateAll(params []{{.ModelName}}Params) error {
		{{toSnakeCase .ModelName}}s := make([]*{{.ModelName}}, len(params))
		for i, p := range params {
			{{toSnakeCase .ModelName}}s[i] = u.Build(p)
		}
		return u.newRelation().CreateAll({{toSnakeCase .ModelName}}s)
	}

	func (r *{{.ModelName}}Relation) CreateAll({{toSnakeCase .ModelName}}s []*{{.ModelName}}) error {
		fieldMap := make(map[string][]interface{})
		for _, {{toSnakeCase .ModelName}} := range {{toSnakeCase .ModelName}}s {
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
						fieldMap["{{toSnakeCase $column}}"] = append(fieldMap["{{toSnakeCase $column}}"], {{toSnakeCase $.ModelName}}.{{$column}})
					{{ end -}}
				}
			}
		}
		return r.Relation.CreateAll(fieldMap)
	}

	func (u *{{.ModelName}}) Update(params {{.ModelName}}Params) error {
		{{ range $column := .Columns -}}
		{{ if eq $column "CreatedAt" -}}
		{{ continue }}
		{{ end -}}
		{{ if eq $column "UpdatedAt" -}}
		{{ continue }}
		{{ end -}}
		if !utils.IsZero(params.{{ $column }}) {
			u.{{ $column }} = params.{{ $column }}
		}
		{{ end -}}
		return u.Save()
	}

	func (m *{{.ModelName}}) Save() error {
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

package template

var textBody = `
		{{define "Base"}}
		// Code generated by ayaorm. DO NOT EDIT.
		package main

		import (
			"fmt"
			"github.com/okazaki-kk/ayaorm"
		)

		type {{.modelName}}Relation struct {
			model *{{.modelName}}
			*ayaorm.Relation
		}

		func (m *{{.modelName}}) newRelation() *{{.modelName}}Relation {
			r := &{{.modelName}}Relation{
				m,
				ayaorm.NewRelation(db).SetTable("{{.snakeCaseModelName}}"),
			}
			r.Select(
				{{ range $column := .columns -}}
				"{{ toSnakeCase $column }}",
				{{ end -}}
			)

			return r
		}

		func (m {{.modelName}}) Select(columns ...string) *{{.modelName}}Relation {
			return m.newRelation().Select(columns...)
		}

		func (r *{{.modelName}}Relation) Select(columns ...string) *{{.modelName}}Relation {
			cs := []string{}
			for _, c := range columns {
				if r.model.isColumnName(c) {
					cs = append(cs, fmt.Sprintf("{{.snakeCaseModelName}}.%s", c))
				} else {
					cs = append(cs, c)
				}
			}
			r.Relation.SetColumns(cs...)
			return r
		}

		type {{.modelName}}Params {{.modelName}}

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

		func (r *{{.modelName}}Relation) QueryRow() (*{{.modelName}}, error) {
			row := &{{.modelName}}{}
			err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
			if err != nil {
				return nil, err
			}
			return row, nil
		}

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

		func (r *{{.modelName}}Relation) Query() ([]*{{.modelName}}, error) {
			rows, err := r.Relation.Query()
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			results := []*{{.modelName}}{}
			for rows.Next() {
				row := &{{.modelName}}{}
				err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
				if err != nil {
					return nil, err
				}
				results = append(results, row)
			}
			return results, nil
		}

		func (m *{{.modelName}}) fieldPtrByName(name string) interface{} {
			switch name {
				{{ range $column := .columns -}}
				case "{{ toSnakeCase  $column}}", "{{$.snakeCaseModelName}}.{{toSnakeCase $column}}":
					return &m.{{$column}}
				{{ end -}}
			default:
				return nil
			}
		}

		func (m *{{.modelName}}) fieldPtrsByName(names []string) []interface{} {
			fields := []interface{}{}
			for _, n := range names {
				f := m.fieldPtrByName(n)
				fields = append(fields, f)
			}
			return fields
		}

		func (m *{{.modelName}}) isColumnName(name string) bool {
			for _, c := range m.columnNames() {
				if c == name {
					return true
				}
			}
			return false
		}

		func (m *{{.modelName}}) columnNames() []string {
			return []string{
				{{ range $column := .columns -}}
				"{{ toSnakeCase $column }}",
				{{ end -}}
			}
		}
		{{end}}
	`

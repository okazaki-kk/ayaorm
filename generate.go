package ayaorm

import (
	"os"
	"regexp"
	"strings"
	"text/template"
)

func Generate(modelName string, field map[string]string) {
	textBody := `
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
				ayaorm.NewRelation(db).SetTable("{{toMultipleSnakeCase .modelName}}"),
			}
			r.Select(
				"id",
				"name",
				"age",
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
					cs = append(cs, fmt.Sprintf("{{toMultipleSnakeCase .modelName}}.%s", c))
				} else {
					cs = append(cs, c)
				}
			}
			r.Relation.SetColumns(cs...)
			return r
		}

		type UserParams {{.modelName}}

		func (m {{.modelName}}) Build(p UserParams) *{{.modelName}} {
			return &{{.modelName}}{
				Id:   p.Id,
				Name: p.Name,
				Age:  p.Age,
			}
		}

		func (r *{{.modelName}}Relation) QueryRow() (*{{.modelName}}, error) {
			row := &{{.modelName}}{}
			fmt.Println(r.Relation.GetColumns())
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

		func (m {{.modelName}}) Save() error {
			return m.newRelation().Save()
		}

		func (r *{{.modelName}}Relation) Save() error {
			fieldMap := make(map[string]interface{})
			for _, c := range r.Relation.GetColumns() {
				switch c {
				case "id", "{{toMultipleSnakeCase .modelName}}.id":
					fieldMap["id"] = r.model.Id
				case "name", "{{toMultipleSnakeCase .modelName}}.name":
					fieldMap["name"] = r.model.Name
				case "age", "{{toMultipleSnakeCase .modelName}}.age":
					fieldMap["age"] = r.model.Age
				}
			}

			return r.Relation.Save(fieldMap)
		}

		func (m *{{.modelName}}) fieldPtrByName(name string) interface{} {
			switch name {
			case "id", "{{toMultipleSnakeCase .modelName}}.id":
				return &m.Id
			case "name", "{{toMultipleSnakeCase .modelName}}.name":
				return &m.Name
			case "age", "{{toMultipleSnakeCase .modelName}}.age":
				return &m.Age
			default:
				return nil
			}
		}

		func (m *{{.modelName}}) fieldPtrsByName(names []string) []interface{} {
			fields := []interface{}{}
			for _, n := range names {
				f := m.fieldPtrByName(n)
				fields = append(fields, f)
				fmt.Println(&f)
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
				"id",
				"name",
				"age",
			}
		}
		{{end}}
	`

	funcMap := template.FuncMap{
		"toMultipleSnakeCase": toMultipleSnakeCase,
	}
	t, _ := template.New("Base").Funcs(funcMap).Parse(textBody)
	f, _ := os.Create("./main_gen.go")
	defer f.Close()
	tmp := make(map[string]string)
	tmp["modelName"] = modelName
	t.Execute(f, tmp)
}

func toMultipleSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake)) + "s"
}

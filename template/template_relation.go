package template

var relationTextBody = `
		{{define "Relation"}}
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
		{{end}}
`

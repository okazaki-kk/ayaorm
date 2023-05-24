package templates

var RelationTextBody = `
		{{define "Relation"}}
		type {{.ModelName}}Relation struct {
			model *{{.ModelName}}
			*ayaorm.Relation
		}

		func (m *{{.ModelName}}) newRelation() *{{.ModelName}}Relation {
			r := &{{.ModelName}}Relation{
				m,
				ayaorm.NewRelation(db).SetTable("{{.SnakeCaseModelName}}"),
			}
			r.Select(
				{{ range $column := .Columns -}}
				"{{ toSnakeCase $column }}",
				{{ end -}}
			)

			return r
		}

		func (m {{.ModelName}}) Select(columns ...string) *{{.ModelName}}Relation {
			return m.newRelation().Select(columns...)
		}

		func (r *{{.ModelName}}Relation) Select(columns ...string) *{{.ModelName}}Relation {
			cs := []string{}
			for _, c := range columns {
				if r.model.isColumnName(c) {
					cs = append(cs, fmt.Sprintf("{{.SnakeCaseModelName}}.%s", c))
				} else {
					cs = append(cs, c)
				}
			}
			r.Relation.SetColumns(cs...)
			return r
		}

		type {{.ModelName}}Params {{.ModelName}}
		{{end}}
`

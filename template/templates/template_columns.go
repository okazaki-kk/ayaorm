package templates

var ColumnsTextBody = `
		{{define "Columns"}}
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

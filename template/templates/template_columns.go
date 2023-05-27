package templates

var ColumnsTextBody = `
		{{define "Columns"}}
		func (m *{{.ModelName}}) fieldPtrByName(name string) interface{} {
			switch name {
				{{ range $column := .Columns -}}
				case "{{ toSnakeCase  $column}}", "{{$.SnakeCaseModelName}}.{{toSnakeCase $column}}":
					return &m.{{$column}}
				{{ end -}}
			default:
				return nil
			}
		}

		func (m *{{.ModelName}}) fieldValuesByName(name string) interface{} {
			switch name {
				{{ range $column := .Columns -}}
				case "{{ toSnakeCase  $column}}", "{{$.SnakeCaseModelName}}.{{toSnakeCase $column}}":
					return m.{{$column}}
				{{ end -}}
			default:
				return nil
			}
		}

		func (m *{{.ModelName}}) fieldPtrsByName(names []string) []interface{} {
			fields := []interface{}{}
			for _, n := range names {
				f := m.fieldPtrByName(n)
				fields = append(fields, f)
			}
			return fields
		}

		func (m *{{.ModelName}}) isColumnName(name string) bool {
			for _, c := range m.columnNames() {
				if c == name {
					return true
				}
			}
			return false
		}

		func (m *{{.ModelName}}) columnNames() []string {
			return []string{
				{{ range $column := .Columns -}}
				"{{ toSnakeCase $column }}",
				{{ end -}}
			}
		}
		{{end}}

`

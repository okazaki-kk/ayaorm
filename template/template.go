package template

var TextBody = `
		{{define "Base"}}
		{{ template "Relation" . }}

		{{ template "CRUD" . }}

		{{ template "Search" . }}

		{{ template "Query" . }}

		{{ template "Columns" . }}
		{{end}}
	`

var FuncBody = `
		{{define "Base"}}
		{{ template "Joins" . }}
		{{end}}
`

var dbTextBody = `
		import "database/sql"

		var db *sql.DB
`

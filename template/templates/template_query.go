package templates

var QueryTextBody = `
	{{define "Query"}}
	func (r *{{.ModelName}}Relation) Query() ([]*{{.ModelName}}, error) {
		rows, err := r.Relation.Query()
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		results := []*{{.ModelName}}{}
		for rows.Next() {
			row := &{{.ModelName}}{}
			err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
			if err != nil {
				return nil, err
			}
			results = append(results, row)
		}
		return results, nil
	}

	func (r *{{.ModelName}}Relation) QueryRow() (*{{.ModelName}}, error) {
		row := &{{.ModelName}}{}
		err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		return row, nil
	}
	{{end}}
`

package templates

var QueryTextBody = `
	{{define "Query"}}
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

	func (r *{{.modelName}}Relation) QueryRow() (*{{.modelName}}, error) {
		row := &{{.modelName}}{}
		err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		return row, nil
	}
	{{end}}
`

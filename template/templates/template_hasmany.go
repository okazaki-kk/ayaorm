package templates

var HasManyTextBody = `func (m {{.Recv}}) {{.HasManyModel}}s() ([]*{{.HasManyModel}}, error) {
	m.hasMany{{.HasManyModel}}s()
	c, err := {{.HasManyModel}}{}.Where("{{toSnakeCase .Recv}}_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (m *{{.Recv}}) DeleteDependent() error {
	{{toSnakeCase .HasManyModel}}s, err := m.{{.HasManyModel}}s()
	if err != nil {
		return err
	}
	for _, {{toSnakeCase .HasManyModel}} := range {{toSnakeCase .HasManyModel}}s {
		err := {{toSnakeCase .HasManyModel}}.Delete()
		if err != nil {
			return err
		}
	}
	err = m.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u {{.Recv}}) Join{{.HasManyModel}}s() *{{.Recv}}Relation {
	return u.newRelation().Join{{.HasManyModel}}s()
}

func (u *{{.Recv}}Relation) Join{{.HasManyModel}}s() *{{.Recv}}Relation {
	u.Relation.InnerJoin("{{toSnakeCase .Recv}}s", "{{toSnakeCase .HasManyModel}}s", true)
	return u
}
`

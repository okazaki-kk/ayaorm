package templates

var HasManyTextBody = `func (m {{.Recv}}) {{.HasManyModel}}s() ([]*{{.HasManyModel}}, error) {
	m.hasMany{{.HasManyModel}}s()
	c, err := {{.HasManyModel}}{}.Where("{{toSnakeCase .Recv}}_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (u {{.Recv}}) Join{{.HasManyModel}}s() *{{.Recv}}Relation {
	return u.newRelation().Join{{.HasManyModel}}s()
}

func (u *{{.Recv}}Relation) Join{{.HasManyModel}}s() *{{.Recv}}Relation {
	u.Relation.InnerJoin("{{toSnakeCase .Recv}}s", "{{toSnakeCase .HasManyModel}}s", true)
	return u
}
`

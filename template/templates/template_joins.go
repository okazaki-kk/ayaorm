package templates

var JoinsTextBody = `func (m {{.Recv}}) {{.HasManyModel}}s() ([]*{{.HasManyModel}}, error) {
	c, err := {{.HasManyModel}}{}.Where("{{toSnakeCase .HasManyModel}}_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (u {{.Recv}}) Join{{.HasManyModel}}s() *{{.Recv}}Relation {
	return u.newRelation().Join{{.HasManyModel}}s()
}

func (u *{{.Recv}}Relation) Join{{.HasManyModel}}s() *{{.Recv}}Relation {
	u.Relation.InnerJoin("posts", "comments")
	return u
}
`

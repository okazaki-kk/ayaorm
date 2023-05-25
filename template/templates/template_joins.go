package templates

var JoinsTextBody = `func (m {{.HasManyModel}}) {{.Recv}}s() ([]*{{.Recv}}, error) {
	c, err := {{.Recv}}{}.Where("{{toSnakeCase .HasManyModel}}_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (u {{.HasManyModel}}) Join{{.Recv}}s() *{{.HasManyModel}}Relation {
	return u.newRelation().Join{{.Recv}}s()
}

func (u *{{.HasManyModel}}Relation) Join{{.Recv}}s() *{{.HasManyModel}}Relation {
	u.Relation.InnerJoin("posts", "comments")
	return u
}
`

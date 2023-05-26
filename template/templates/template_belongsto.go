package templates

var BelongsTextBody = `func (u {{.Recv}}) {{.BelongsToModel}}() (*{{.BelongsToModel}}, error) {
	return {{.BelongsToModel}}{}.Find(u.{{.BelongsToModel}}Id)
}

func (u {{.Recv}}) Join{{.BelongsToModel}}() *{{.Recv}}Relation {
	return u.newRelation().Join{{.BelongsToModel}}()
}

func (u *{{.Recv}}Relation) Join{{.BelongsToModel}}() *{{.Recv}}Relation {
	u.Relation.InnerJoin("{{toSnakeCase .Recv}}s", "{{toSnakeCase .BelongsToModel}}s", false)
	return u
}
`

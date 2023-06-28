package templates

var HasOneTextBody = `func (u {{.Recv}}) {{.HasOneModel}}()  (*{{.HasOneModel}}, error) {
	u.hasOne{{.HasOneModel}}()
	c, err := {{.HasOneModel}}{}.FindBy("{{toSnakeCase .Recv}}_id", u.Id)
	if err != nil {
		return nil, err
	}
	return c, nil
}
`

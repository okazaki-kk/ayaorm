package templates

var HasOneTextBody = `func (u {{.Recv}}) {{.HasOneModel}}()  (*{{.HasOneModel}}, error) {
	u.hasOne{{.HasOneModel}}()
	return {{.HasOneModel}}{}.Find(u.Id)
}
`

package templates

var HasOneTextBody = `func (u Post) Project()  (*Project, error) {
	u.hasOneProject()
	return Project{}.Find(u.Id)
}
`

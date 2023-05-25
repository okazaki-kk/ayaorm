package templates

var JoinsTextBody = `func (m {{.Recv}}) Comments() ([]*Comment, error) {
	comments, err := Comment{}.Where("post_id", m.Id).Query()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (u {{.Recv}}) JoinComments() *{{.Recv}}Relation {
	return u.newRelation().JoinComments()
}

func (u *{{.Recv}}Relation) JoinComments() *{{.Recv}}Relation {
	u.Relation.InnerJoin("posts", "comments")
	return u
}
`

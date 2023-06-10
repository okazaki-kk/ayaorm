package query

type innerJoin struct {
	left    string
	right   string
	hasMany bool
}

func (q *Query) InnerJoin(left, right string, hasMany bool) *Query {
	q.innerJoin.left = left
	q.innerJoin.right = right
	q.innerJoin.hasMany = hasMany
	return q
}

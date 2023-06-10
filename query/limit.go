package query

func (q *Query) Limit(lim int) *Query {
	q.limit = lim
	return q
}

func (q *Query) Order(key, value string) *Query {
	q.order = value
	q.orderKey = key
	return q
}

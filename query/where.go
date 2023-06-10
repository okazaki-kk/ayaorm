package query

type where struct {
	key        string
	conditions []interface{}
}

func (q *Query) Where(column string, conditions ...interface{}) *Query {
	q.where.key = column
	q.where.conditions = conditions
	return q
}

type or struct {
	key        string
	conditions []interface{}
}

func (q *Query) Or(column string, conditions ...interface{}) *Query {
	q.or.key = column
	q.or.conditions = conditions
	return q
}

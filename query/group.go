package query

type having struct {
	key        string
	conditions []interface{}
}

func (q *Query) Having(column string, conditions ...interface{}) *Query {
	q.having.key = column
	q.having.conditions = conditions
	return q
}

func (q *Query) GroupBy(column ...string) *Query {
	q.groupBy = column
	return q
}

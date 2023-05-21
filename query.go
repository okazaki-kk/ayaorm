package ayaorm

import (
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	limit     int
	order     string
	orderKey  string
	where     struct {
		key   string
		value interface{}
	}
	insert    struct{ params map[string]interface{} }
	update    struct{ params map[string]interface{} }
	innerJoin struct {
		left  string
		right string
	}
}

func (q *Query) BuildQuery(columns []string, tableName string) string {
	query := fmt.Sprintf("SELECT % s FROM %s", strings.Join(columns, ", "), tableName)
	if q.where.key != "" {
		if reflect.TypeOf(q.where.value).Kind() == reflect.String {
			query = fmt.Sprintf("%s WHERE %s = '%s'", query, q.where.key, q.where.value)
		} else {
			query = fmt.Sprintf("%s WHERE %s = %d", query, q.where.key, q.where.value)
		}
	}
	if q.order != "" {
		query = fmt.Sprintf("%s ORDER BY %s %s", query, q.orderKey, q.order)
	}
	if q.limit > 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, q.limit)
	}
	if q.innerJoin.left != "" {
		text := q.innerJoin.left[:len(q.innerJoin.left)-1]
		query = fmt.Sprintf("%s INNER JOIN %s on %s.id = %s.%s_id", query, q.innerJoin.right, q.innerJoin.left, q.innerJoin.right, text)
	}
	return query + ";"
}
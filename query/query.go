package query

import (
	"fmt"
	"strings"
)

type Query struct {
	limit    int
	order    string
	orderKey string
	where
	or
	groupBy []string
	having
	insert
	update
	createAll
	innerJoin
}

func (q *Query) BuildQuery(columns []string, tableName string) (string, []interface{}) {
	query := fmt.Sprintf("SELECT % s FROM %s", strings.Join(columns, ", "), tableName)
	var args []interface{}

	if q.where.key != "" {
		switch len(q.where.conditions) {
		case 1:
			if q.where.conditions[0] == nil {
				query = fmt.Sprintf("%s WHERE %s IS NULL", query, q.where.key)
			} else {
				query = fmt.Sprintf("%s WHERE %s = ?", query, q.where.key)
				args = append(args, q.where.conditions[0])
			}
		case 2:
			query = fmt.Sprintf("%s WHERE %s %s ?", query, q.where.key, q.where.conditions[0])
			args = append(args, q.where.conditions[1])
		case 3:
			query = fmt.Sprintf("%s WHERE %s %s ?", query, q.where.key, q.where.conditions[0])
		default:
			query = ""
		}
	}

	if q.or.key != "" {
		switch len(q.or.conditions) {
		case 1:
			if q.or.conditions[0] == nil {
				query = fmt.Sprintf("%s OR %s IS NULL", query, q.or.key)
			} else {
				query = fmt.Sprintf("%s OR %s = ?", query, q.or.key)
				args = append(args, q.or.conditions[0])
			}
		case 2:
			query = fmt.Sprintf("%s OR %s %s ?", query, q.or.key, q.or.conditions[0])
			args = append(args, q.or.conditions[1])
		case 3:
			query = fmt.Sprintf("%s OR %s %s ?", query, q.or.key, q.or.conditions[0])
		default:
			query = ""
		}
	}

	if q.groupBy != nil {
		query = fmt.Sprintf("%s GROUP BY %s", query, strings.Join(q.groupBy, ", "))
	}

	if q.having.key != "" {
		switch len(q.having.conditions) {
		case 1:
			query = fmt.Sprintf("%s HAVING %s = ?", query, q.having.key)
			args = append(args, q.having.conditions[0])
		case 2:
			query = fmt.Sprintf("%s HAVING %s %s ?", query, q.having.key, q.having.conditions[0])
			args = append(args, q.having.conditions[1])
		case 3:
			query = fmt.Sprintf("%s HAVING %s %s ?", query, q.having.key, q.having.conditions[0])
		default:
			query = ""
		}
	}

	if q.order != "" {
		query = fmt.Sprintf("%s ORDER BY %s %s", query, q.orderKey, q.order)
	}

	if q.limit > 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, q.limit)
	}

	if q.innerJoin.left != "" {
		textLeft := q.innerJoin.left[:len(q.innerJoin.left)-1]
		textRight := q.innerJoin.right[:len(q.innerJoin.right)-1]

		if q.innerJoin.hasMany {
			query = fmt.Sprintf("%s INNER JOIN %s on %s.id = %s.%s_id", query, q.innerJoin.right, q.innerJoin.left, q.innerJoin.right, textLeft)
		} else {
			query = fmt.Sprintf("%s INNER JOIN %s on %s.%s_id = %s.id", query, q.innerJoin.right, q.innerJoin.left, textRight, q.innerJoin.right)
		}
	}

	return query + ";", args
}

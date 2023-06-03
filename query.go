package ayaorm

import (
	"fmt"
	"strings"
)

type Query struct {
	limit    int
	order    string
	orderKey string
	where    struct {
		key        string
		conditions []interface{}
	}
	or struct {
		key        string
		conditions []interface{}
	}
	groupBy   []string
	insert    struct{ params map[string]interface{} }
	update    struct{ params map[string]interface{} }
	innerJoin struct {
		left    string
		right   string
		hasMany bool
	}
}

func (q *Query) Where(column string, conditions ...interface{}) *Query {
	q.where.key = column
	q.where.conditions = conditions
	return q
}

func (q *Query) Or(column string, conditions ...interface{}) *Query {
	q.or.key = column
	q.or.conditions = conditions
	return q
}

func (q *Query) InnerJoin(left, right string, hasMany bool) *Query {
	q.innerJoin.left = left
	q.innerJoin.right = right
	q.innerJoin.hasMany = hasMany
	return q
}

func (q *Query) BuildQuery(columns []string, tableName string) (string, []interface{}) {
	query := fmt.Sprintf("SELECT % s FROM %s", strings.Join(columns, ", "), tableName)
	var args []interface{}

	if q.where.key != "" {
		switch len(q.where.conditions) {
		case 1:
			query = fmt.Sprintf("%s WHERE %s = ?", query, q.where.key)
			args = append(args, q.where.conditions[0])
		case 2:
			query = fmt.Sprintf("%s WHERE %s %s ?", query, q.where.key, q.where.conditions[0])
			args = append(args, q.where.conditions[1])
		case 3:
			query = fmt.Sprintf("%s WHERE %s %s ?", query, q.where.key, q.where.conditions[0])
		default:
			query = ""
		}
	}

	if q.groupBy != nil {
		query = fmt.Sprintf("%s GROUP BY %s", query, strings.Join(q.groupBy, ", "))
	}

	if q.or.key != "" {
		switch len(q.or.conditions) {
		case 1:
			query = fmt.Sprintf("%s OR %s = ?", query, q.or.key)
			args = append(args, q.or.conditions[0])
		case 2:
			query = fmt.Sprintf("%s OR %s %s ?", query, q.or.key, q.or.conditions[0])
			args = append(args, q.or.conditions[1])
		case 3:
			query = fmt.Sprintf("%s OR %s %s ?", query, q.or.key, q.or.conditions[0])
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

func (q *Query) BuildInsert(tableName string) (string, []interface{}) {
	columns := []string{}
	ph := []string{}
	args := []interface{}{}
	i := q.insert

	for k, v := range i.params {
		columns = append(columns, k)
		ph = append(ph, "?")
		args = append(args, v)
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(ph, ", ")), args
}

func (q *Query) BuildUpdate(tableName string, id int) (string, []interface{}) {
	args := []interface{}{}
	i := q.update

	updateObj := ""

	for k, v := range i.params {
		updateObj = fmt.Sprintf("%s %s = ?,", updateObj, k)
		args = append(args, v)
	}
	updateObj = updateObj[:len(updateObj)-1]
	fmt.Println(updateObj, "update")

	return fmt.Sprintf("UPDATE %s SET %s WHERE id = %d;", tableName, updateObj, id), args
}

func (q *Query) BuildDelete(tableName string, id int) string {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", tableName, id)
	return query
}

func (q *Query) BuildInnerJoin(left, right string) string {
	query := fmt.Sprintf("SELECT * FROM %s inner join %s on %s.id = %s.post_id", left, right, left, right)
	return query
}

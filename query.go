package ayaorm

import (
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	limit    int
	order    string
	orderKey string
	where    struct {
		key   string
		value interface{}
	}
	insert    struct{ params map[string]interface{} }
	update    struct{ params map[string]interface{} }
	innerJoin struct {
		left    string
		right   string
		hasMany bool
	}
}

func (q *Query) Where(column string, value interface{}) *Query {
	q.where.key = column
	q.where.value = value
	return q
}

func (q *Query) InnerJoin(left, right string, hasMany bool) *Query {
	q.innerJoin.left = left
	q.innerJoin.right = right
	q.innerJoin.hasMany = hasMany
	return q
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
		textLeft := q.innerJoin.left[:len(q.innerJoin.left)-1]
		textRight := q.innerJoin.right[:len(q.innerJoin.right)-1]

		if q.innerJoin.hasMany {
			query = fmt.Sprintf("%s INNER JOIN %s on %s.id = %s.%s_id", query, q.innerJoin.right, q.innerJoin.left, q.innerJoin.right, textLeft)
		} else {
			query = fmt.Sprintf("%s INNER JOIN %s on %s.%s_id = %s.id", query, q.innerJoin.right, q.innerJoin.left, textRight, q.innerJoin.right)
		}
	}
	return query + ";"
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

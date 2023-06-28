package query

import (
	"fmt"
	"strings"

	"github.com/okazaki-kk/ayaorm/null"
)

type insert struct {
	params map[string]interface{}
}

type update struct {
	params map[string]interface{}
}

type createAll struct {
	params map[string][]interface{}
}

func (q *Query) Insert(params map[string]interface{}) *Query {
	q.insert.params = params
	return q
}

func (q *Query) Update(params map[string]interface{}) *Query {
	q.update.params = params
	return q
}

func (q *Query) CreateAll(params map[string][]interface{}) *Query {
	q.createAll.params = params
	return q
}

func (q *Query) BuildInsert(tableName string) (string, []interface{}) {
	columns := []string{}
	ph := []string{}
	args := []interface{}{}
	i := q.insert

	for k, v := range i.params {
		if nu, ok := v.(null.Null); ok && !nu.Valid() {
			continue
		}

		columns = append(columns, k)
		ph = append(ph, "?")
		args = append(args, v)
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(ph, ", ")), args
}

func (q *Query) BuildCreateAll(tableName string) (string, []interface{}) {
	columns := []string{}
	ph := ""
	args := []interface{}{}

	for column := range q.createAll.params {
		columns = append(columns, column)
	}

	xCount := len(q.createAll.params[columns[0]])
	yCount := len(columns)

	for i := 0; i < xCount; i++ {
		ph += "("
		for j := 0; j < yCount; j++ {
			if j == yCount-1 {
				ph += "?"
			} else {
				ph += "?, "
			}
			args = append(args, q.createAll.params[columns[j]][i])
		}
		if i == xCount-1 {
			ph += ")"
		} else {
			ph += "), "
		}
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", tableName, strings.Join(columns, ", "), ph), args
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

	return fmt.Sprintf("UPDATE %s SET %s WHERE id = %d;", tableName, updateObj, id), args
}

func (q *Query) BuildDelete(tableName string, id int) string {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", tableName, id)
	return query
}

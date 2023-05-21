package ayaorm

import (
	"fmt"
	"strings"
)

type Table struct {
	tableName string
	columns   []string
	query     Query
}

func (s *Table) SetTable(tableName string) *Table {
	s.tableName = tableName
	return s
}

func (s *Table) SetColumns(columns ...string) *Table {
	s.columns = columns
	return s
}

func (s *Table) GetColumns() []string {
	return s.columns
}

func (s *Table) Where(column string, value interface{}) *Table {
	s.query.where.key = column
	s.query.where.value = value
	return s
}

func (s *Table) InnerJoin(left, right string) *Table {
	s.query.innerJoin.left = left
	s.query.innerJoin.right = right
	return s
}

func (s *Table) BuildInsert() (string, []interface{}) {
	columns := []string{}
	ph := []string{}
	args := []interface{}{}
	i := s.query.insert

	for k, v := range i.params {
		columns = append(columns, k)
		ph = append(ph, "?")
		args = append(args, v)
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", s.tableName, strings.Join(columns, ", "), strings.Join(ph, ", ")), args
}

func (s *Table) BuildUpdate(id int) (string, []interface{}) {
	args := []interface{}{}
	i := s.query.update

	updateObj := ""

	for k, v := range i.params {
		updateObj = fmt.Sprintf("%s %s = ?,", updateObj, k)
		args = append(args, v)
	}
	updateObj = updateObj[:len(updateObj)-1]
	fmt.Println(updateObj, "update")

	return fmt.Sprintf("UPDATE %s SET %s WHERE id = %d;", s.tableName, updateObj, id), args
}

func (s *Table) BuildDelete(id int) string {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", s.tableName, id)
	return query
}

func (s *Table) BuildInnerJoin(left, right string) string {
	query := fmt.Sprintf("SELECT * FROM %s inner join %s on %s.id = %s.post_id", left, right, left, right)
	return query
}

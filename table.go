package ayaorm

import (
	"fmt"
	"reflect"
	"strings"
)

type Table struct {
	tableName string
	columns   []string
	limit     int
	order     string
	orderKey  string
	where     struct {
		key   string
		value interface{}
	}
	insert struct{ params map[string]interface{} }
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
	s.where.key = column
	s.where.value = value
	return s
}

func (s *Table) BuildInsert() (string, []interface{}) {
	columns := []string{}
	ph := []string{}
	args := []interface{}{}
	i := s.insert

	for k, v := range i.params {
		columns = append(columns, k)
		ph = append(ph, "?")
		args = append(args, v)
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", s.tableName, strings.Join(columns, ", "), strings.Join(ph, ", ")), args
}

func (s *Table) BuildQuery() string {
	query := fmt.Sprintf("SELECT % s FROM %s", strings.Join(s.columns, ", "), s.tableName)
	if s.order != "" {
		query = fmt.Sprintf("%s ORDER BY %s %s", query, s.orderKey, s.order)
	}
	if s.limit > 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, s.limit)
	}
	if s.where.key != "" {
		if reflect.TypeOf(s.where.value).Kind() == reflect.String {
			query = fmt.Sprintf("%s WHERE %s = '%s'", query, s.where.key, s.where.value)
		} else {
			query = fmt.Sprintf("%s WHERE %s = %d", query, s.where.key, s.where.value)
		}
	}
	return query + ";"
}

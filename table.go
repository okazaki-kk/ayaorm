package ayaorm

import (
	"fmt"
	"strings"
)

type Table struct {
	tableName string
	columns   []string
	limit     int
	order     string
	orderKey  string
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

func (s *Table) BuildQuery() string {
	query := fmt.Sprintf("SELECT % s FROM %s", strings.Join(s.columns, ", "), s.tableName)
	if s.limit > 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, s.limit)
	}
	if s.order != "" {
		query = fmt.Sprintf("%s ORDER BY %s %s", query, s.orderKey, s.order)
	}
	return query + ";"
}

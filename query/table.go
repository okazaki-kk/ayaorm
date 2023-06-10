package query

type Table struct {
	tableName string
	columns   []string
	Query
}

func (s *Table) SetTable(tableName string) *Table {
	s.tableName = tableName
	return s
}

func (s *Table) GetTable() string {
	return s.tableName
}

func (s *Table) SetColumns(columns ...string) *Table {
	s.columns = columns
	return s
}

func (s *Table) GetColumns() []string {
	return s.columns
}

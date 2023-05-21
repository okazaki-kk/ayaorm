package ayaorm

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

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

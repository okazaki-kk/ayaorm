package ayaorm

import (
	"database/sql"
	"log"
)

type Relation struct {
	Table
	db *sql.DB
}

func NewRelation(db *sql.DB) *Relation {
	r := &Relation{db: db, Table: Table{}}
	return r
}

func (r *Relation) SetTable(tableName string) *Relation {
	r.Table.SetTable(tableName)
	return r
}

func (r *Relation) SetColumns(columns ...string) *Relation {
	r.Table.SetColumns(columns...)
	return r
}

func (r *Relation) Count(column ...string) int {
	var count int
	if err := r.SetColumns("COUNT(*)").QueryRow(&count); err != nil {
		log.Print("ERROR IN COUNTR", err)
		return 0
	}
	return count
}

func (r *Relation) QueryRow(dest ...interface{}) error {
	query := r.BuildQuery()
	log.Print("excute query: ", query)
	return r.db.QueryRow(query).Scan(dest...)
}

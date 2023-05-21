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
		log.Print("ERROR IN COUNTER ", err)
		return 0
	}
	return count
}

func (r *Relation) Limit(lim int) *Relation {
	r.Table.limit = lim
	return r
}

func (r *Relation) Order(key, order string) *Relation {
	r.Table.order = order
	r.Table.orderKey = key
	return r
}

func (r *Relation) Where(column string, value interface{}) *Relation {
	r.Table.Where(column, value)
	return r
}

func (r *Relation) Save(fieldMap map[string]interface{}) (int, error) {
	r.Table.insert.params = fieldMap
	query, args := r.BuildInsert()
	log.Print("excute query: ", query, args)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), err
}

func (r *Relation) Update(id int, fieldMap map[string]interface{}) error {
	r.Table.update.params = fieldMap
	query, args := r.BuildUpdate(id)
	log.Print("excute query: ", query, args)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Relation) Delete(id int) error {
	query := r.BuildDelete(id)
	log.Print("excute query: ", query)
	_, err := r.db.Exec(query)
	return err
}

func (r *Relation) First() *Relation {
	r.Limit(1).Order("id", "asc")
	return r
}

func (r *Relation) Last() *Relation {
	r.Limit(1).Order("id", "desc")
	return r
}

func (r *Relation) Find(id int) *Relation {
	r.Where("id", id)
	r.Limit(1)
	return r
}

func (r *Relation) FindBy(column string, value interface{}) *Relation {
	r.Where(column, value)
	r.Limit(1)
	return r
}

func (r *Relation) InnerJoin(left, right string) *Relation {
	r.Table.InnerJoin(left, right)
	return r
}

func (r *Relation) QueryRow(dest ...interface{}) error {
	query := r.BuildQuery()
	log.Print("excute query: ", query)
	return r.db.QueryRow(query).Scan(dest...)
}

func (r *Relation) Query() (*sql.Rows, error) {
	query := r.BuildQuery()
	log.Print("excute query: ", query)
	rows, err := r.db.Query(query)
	return rows, err
}

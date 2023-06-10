package ayaorm

import (
	"database/sql"
	"log"

	"github.com/okazaki-kk/ayaorm/query"
)

type Relation struct {
	query.Table
	db *sql.DB
}

func NewRelation(db *sql.DB) *Relation {
	r := &Relation{db: db, Table: query.Table{}}
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

func (r *Relation) Pluck(column string) ([]interface{}, error) {
	var res []interface{}
	rows, err := r.SetColumns(column).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp interface{}
		err := rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (r *Relation) Limit(lim int) *Relation {
	r.Table.Limit(lim)
	return r
}

func (r *Relation) Order(key, order string) *Relation {
	r.Table.Order(key, order)
	return r
}

func (r *Relation) Where(column string, conditions ...interface{}) *Relation {
	r.Table.Where(column, conditions...)
	return r
}

func (r *Relation) Or(column string, conditions ...interface{}) *Relation {
	r.Table.Or(column, conditions...)
	return r
}

func (r *Relation) GroupBy(column ...string) *Relation {
	r.Table.GroupBy(column...)
	return r
}

func (r *Relation) Having(column string, conditions ...interface{}) *Relation {
	r.Table.Having(column, conditions...)
	return r
}

func (r *Relation) Save(id int, fieldMap map[string]interface{}) (int, error) {
	r.Table.Insert(fieldMap)
	var query string
	var args []interface{}

	if IsZero(id) {
		query, args = r.Table.BuildInsert(r.GetTable())
	} else {
		r.Table.Update(fieldMap)
		query, args = r.Table.BuildUpdate(r.Table.GetTable(), id)
	}

	log.Print("execute query: ", query, " ", args)

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

func (r *Relation) CreateAll(fieldMaps map[string][]interface{}) error {
	r.Table.CreateAll(fieldMaps)
	var query string
	var args []interface{}

	query, args = r.Table.BuildCreateAll(r.Table.GetTable())

	log.Println("execute query: ", query, " ", args)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Relation) Delete(id int) error {
	query := r.Table.BuildDelete(r.Table.GetTable(), id)
	log.Print("execute query: ", query)
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

func (r *Relation) InnerJoin(left, right string, hasMany bool) *Relation {
	r.Table.InnerJoin(left, right, hasMany)
	return r
}

func (r *Relation) QueryRow(dest ...interface{}) error {
	query, args := r.Table.BuildQuery(r.Table.GetColumns(), r.GetTable())
	log.Print("execute query: ", query, " ", args)
	return r.db.QueryRow(query, args...).Scan(dest...)
}

func (r *Relation) Query() (*sql.Rows, error) {
	query, args := r.Table.BuildQuery(r.Table.GetColumns(), r.GetTable())
	log.Print("execute query: ", query, " ", args)
	rows, err := r.db.Query(query, args...)
	return rows, err
}

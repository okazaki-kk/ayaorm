package ayaorm

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestCount(t *testing.T) {
	db, _ := sql.Open("sqlite3", "./ayaorm.db")
	defer db.Close()
	defer os.Remove("./ayaorm.db")
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER)")
	db.Exec("INSERT INTO users (name, age) VALUES ('Hanako', 20)")
	db.Exec("INSERT INTO users (name, age) VALUES ('Hanako', 20)")
	db.Exec("INSERT INTO users (name, age) VALUES ('Hanako', 20)")

	count := relation.Count()
	assert.Equal(t, count, 3)
}

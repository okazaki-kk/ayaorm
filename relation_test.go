package ayaorm

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type TestUser struct {
	Id   int `db:"pk"`
	Name string
	Age  int
}

func TestMain(m *testing.M) {
	db, _ = sql.Open("sqlite3", "./ayaorm.db")
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		os.Exit(1)
	}
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER)")
	if err != nil {
		os.Exit(1)
	}
	_, err = db.Exec("INSERT INTO users (name, age) VALUES ('Hanako', 20)")
	if err != nil {
		os.Exit(1)
	}
	_, err = db.Exec("INSERT INTO users (name, age) VALUES ('Taro', 23)")
	if err != nil {
		os.Exit(1)
	}

	code := m.Run()
	defer db.Close()
	defer os.Remove("./ayaorm.db")
	os.Exit(code)
}

func TestCount(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	count := relation.Count()
	assert.Equal(t, count, 2)
}

func TestSave(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	err := relation.Save(map[string]interface{}{"name": "Jiro", "age": 25})
	assert.NoError(t, err)

	count := relation.Count()
	assert.Equal(t, count, 3)
}

func TestWhere(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	Hanako, err := relation.SetColumns("*").Where("name", "Hanako").Query()
	assert.NoError(t, err)
	for Hanako.Next() {
		var user TestUser
		err := Hanako.Scan(&user.Id, &user.Name, &user.Age)
		assert.NoError(t, err)
		assert.Equal(t, user.Name, "Hanako")
		assert.Equal(t, user.Age, 20)
	}
	assert.NoError(t, err)
}

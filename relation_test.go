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

func TestSave(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	countBefore := relation.Count()

	id, err := relation.Save(map[string]interface{}{"name": "Jiro", "age": 25})
	assert.NoError(t, err)
	assert.NotZero(t, id)

	countAfter := relation.Count()
	var lastUser TestUser
	err = relation.SetColumns("*").Last().QueryRow(&lastUser.Id, &lastUser.Name, &lastUser.Age)

	assert.NoError(t, err)
	assert.Equal(t, 25, lastUser.Age)
	assert.Equal(t, "Jiro", lastUser.Name)
	assert.Equal(t, countBefore+1, countAfter)
}

func TestUpdate(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	var user TestUser
	err := relation.SetColumns("*").Last().QueryRow(&user.Id, &user.Name, &user.Age)
	assert.NoError(t, err)

	id := user.Id
	age := user.Age

	fieldMap := make(map[string]interface{})
	fieldMap["Name"] = "DigDag"

	relation.Update(id, fieldMap)

	err = relation.SetColumns("*").Last().QueryRow(&user.Id, &user.Name, &user.Age)
	assert.NoError(t, err)
	assert.Equal(t, age, user.Age)
	assert.Equal(t, "DigDag", user.Name)
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
		assert.Equal(t, "Hanako", user.Name)
		assert.Equal(t, 20, user.Age)
	}
	assert.NoError(t, err)
}

func TestFirst(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	var user TestUser
	err := relation.SetColumns("*").First().QueryRow(&user.Id, &user.Name, &user.Age)
	assert.NoError(t, err)
	assert.Equal(t, 20, user.Age)
	assert.Equal(t, "Hanako", user.Name)
}

func TestFind(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	var user TestUser
	err := relation.SetColumns("*").Find(2).QueryRow(&user.Id, &user.Name, &user.Age)
	assert.NoError(t, err)
	assert.Equal(t, 23, user.Age)
	assert.Equal(t, "Taro", user.Name)
}

func TestFindBy(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	var user TestUser
	err := relation.SetColumns("*").FindBy("Name", "Taro").QueryRow(&user.Id, &user.Name, &user.Age)
	assert.NoError(t, err)
	assert.Equal(t, 23, user.Age)
	assert.Equal(t, "Taro", user.Name)
}

func TestDelete(t *testing.T) {
	table := Table{tableName: "users"}
	relation := Relation{Table: table, db: db}

	countBefore := relation.Count()

	var user TestUser
	err := relation.SetColumns("*").Last().QueryRow(&user.Id, &user.Name, &user.Age)
	assert.NoError(t, err)

	err = relation.Delete(user.Id)
	assert.NoError(t, err)

	afterCount := relation.Count()

	assert.Equal(t, countBefore-1, afterCount)

}

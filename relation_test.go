package ayaorm

import (
	"database/sql"
	"os"
	"testing"

	"github.com/okazaki-kk/ayaorm/query"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type TestUser struct {
	Schema
	Name string
	Age  int
}

func TestMain(m *testing.M) {
	db, _ = sql.Open("sqlite3", "./ayaorm.db")
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		os.Exit(1)
	}
	_, err = db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')));`)
	if err != nil {
		os.Exit(1)
	}
	_, err = db.Exec(`
		CREATE TRIGGER trigger_test_updated_at AFTER UPDATE ON users
		BEGIN
			UPDATE users SET updated_at = DATETIME('now', 'localtime') WHERE rowid == NEW.rowid;
		END;
	`)
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
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	countBefore := relation.Count()

	id, err := relation.Save(0, map[string]interface{}{"name": "Jiro", "age": 25})
	assert.NoError(t, err)
	assert.NotZero(t, id)

	countAfter := relation.Count()
	var lastUser TestUser
	err = relation.SetColumns("*").Last().QueryRow(&lastUser.Id, &lastUser.Name, &lastUser.Age, &lastUser.CreatedAt, &lastUser.UpdatedAt)

	assert.NoError(t, err)
	assert.Equal(t, 25, lastUser.Age)
	assert.Equal(t, "Jiro", lastUser.Name)
	assert.Equal(t, countBefore+1, countAfter)

	countBefore = relation.Count()
	id, err = relation.Save(lastUser.Id, map[string]interface{}{"name": "Saburo"})
	assert.NoError(t, err)
	assert.NotZero(t, id)

	countAfter = relation.Count()
	err = relation.SetColumns("*").Last().QueryRow(&lastUser.Id, &lastUser.Name, &lastUser.Age, &lastUser.CreatedAt, &lastUser.UpdatedAt)

	assert.NoError(t, err)
	assert.Equal(t, 25, lastUser.Age)
	assert.Equal(t, "Saburo", lastUser.Name)
	assert.Equal(t, countBefore, countAfter)
}

func TestPluck(t *testing.T) {
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	ages, err := relation.Pluck("age")
	assert.NoError(t, err)
	assert.Equal(t, int64(20), ages[0])
	assert.Equal(t, int64(23), ages[1])
}

func TestWhere(t *testing.T) {
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	Hanako, err := relation.SetColumns("*").Where("name", "Hanako").Query()
	assert.NoError(t, err)
	for Hanako.Next() {
		var user TestUser
		err := Hanako.Scan(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		assert.NoError(t, err)
		assert.Equal(t, "Hanako", user.Name)
		assert.Equal(t, 20, user.Age)
	}
}

func TestWhere1(t *testing.T) {
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	Hanako, err := relation.SetColumns("*").Where("age", ">", 24).Query()
	assert.NoError(t, err)
	for Hanako.Next() {
		var user TestUser
		err := Hanako.Scan(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		assert.NoError(t, err)
		assert.Equal(t, "Saburo", user.Name)
		assert.Equal(t, 25, user.Age)
	}
}

func TestFirst(t *testing.T) {
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	var user TestUser
	err := relation.SetColumns("*").First().QueryRow(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	assert.NoError(t, err)
	assert.Equal(t, 20, user.Age)
	assert.Equal(t, "Hanako", user.Name)
}

func TestFind(t *testing.T) {
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	var user TestUser
	err := relation.SetColumns("*").Find(2).QueryRow(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	assert.NoError(t, err)
	assert.Equal(t, 23, user.Age)
	assert.Equal(t, "Taro", user.Name)
}

func TestFindBy(t *testing.T) {
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	var user TestUser
	err := relation.SetColumns("*").FindBy("Name", "Taro").QueryRow(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	assert.NoError(t, err)
	assert.Equal(t, 23, user.Age)
	assert.Equal(t, "Taro", user.Name)
}

func TestDelete(t *testing.T) {
	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	countBefore := relation.Count()

	var user TestUser
	err := relation.SetColumns("*").Last().QueryRow(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	assert.NoError(t, err)

	err = relation.Delete(user.Id)
	assert.NoError(t, err)

	afterCount := relation.Count()

	assert.Equal(t, countBefore-1, afterCount)

}

type TestComment struct {
	Schema
	Content string
	UserId  int
}

func TestInnerJoin(t *testing.T) {
	_, err := db.Exec(`drop table if exists comments`)
	assert.NoError(t, err)

	_, err = db.Exec(`create table comments (id integer primary key autoincrement, content text, user_id integer, created_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now', 'localtime')), updated_at TIMESTAMP NOT NULL DEFAULT(DATETIME('now','localtime')), foreign key (user_id) references users(id) );`)
	assert.NoError(t, err)

	_, err = db.Exec(`insert into comments (content, user_id) values ('Wonderful', 1)`)
	assert.NoError(t, err)

	_, err = db.Exec(`insert into comments (content, user_id) values ('Bad', 1)`)
	assert.NoError(t, err)

	table := query.Table{}
	table.SetTable("users")
	relation := Relation{Table: table, db: db}

	rows, err := relation.SetColumns("users.id, users.name, users.age, users.created_at, users.updated_at").InnerJoin("users", "comments", true).Query()
	assert.NoError(t, err)

	for rows.Next() {
		var user TestUser
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		assert.NoError(t, err)
		assert.Equal(t, user.Id, 1)
		assert.Equal(t, user.Name, "Hanako")
		assert.Equal(t, user.Age, 20)
	}
}

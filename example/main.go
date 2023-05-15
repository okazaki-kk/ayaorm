package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// +AYAORM
type User struct {
	Id   int `db:"pk"`
	Name string
	Age  int
}

func main() {
	db, _ = sql.Open("sqlite3", "./ayaorm.db")

	fmt.Println("USER COUNT:", User{}.Count())

	userRows, err := User{}.All().Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	defer userRows.Close()

	for userRows.Next() {
		var user User
		if err := userRows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Fatal("User.All.Scan.Error", err)
		}
		fmt.Println("User.All.Scan:", user)
	}

	userRows, err = User{}.Limit(2).Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	defer userRows.Close()

	for userRows.Next() {
		var user User
		if err := userRows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Fatal("User.All.Scan.Error", err)
		}
		fmt.Println("User.All.Scan:", user)
	}

	reverseUserRows, err := User{}.Order("Age", "desc").Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	defer reverseUserRows.Close()

	for reverseUserRows.Next() {
		var user User
		if err := reverseUserRows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Fatal("User.All.Scan.Error", err)
		}
		fmt.Println("User.All.Scan:", user)
	}

}

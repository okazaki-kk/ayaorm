package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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

	Hanako, err := User{}.Where("Name", "Hanako").Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	defer Hanako.Close()

	for Hanako.Next() {
		var user User
		if err := Hanako.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Fatal("User.All.Scan.Error", err)
		}
		fmt.Println("Hanako:", user)
	}

	Age34, err := User{}.Where("Age", 34).Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	defer Age34.Close()

	for Age34.Next() {
		var user User
		if err := Age34.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Fatal("User.All.Scan.Error", err)
		}
		fmt.Println("Age34:", user)
	}

	newUser := User{Name: "Gon", Age: 20}
	err = newUser.Save()
	if err != nil {
		log.Fatal("User.Save.Error", err)
	}

}

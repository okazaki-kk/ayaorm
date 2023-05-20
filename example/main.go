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

	users, err := User{}.All().Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}

	users, err = User{}.Limit(2).Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}

	reverseUsers, err := User{}.Order("Age", "desc").Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	for _, user := range reverseUsers {
		fmt.Printf("%v\n", user)
	}

	Hanakos, err := User{}.Where("Name", "Hanako").Query()
	if err != nil {
		log.Fatal("User.All.Error", err)
	}
	for _, hanako := range Hanakos {
		fmt.Printf("%v\n", hanako)
	}

	newUser := User{Name: "Gin", Age: 34}
	err = newUser.Save()
	if err != nil {
		log.Fatal("User.Save.Error", err)
	}
	fmt.Printf("%v\n", newUser)

	firstUser, err := User{}.First()
	if err != nil {
		log.Fatal("User.First.Error", err)
	}
	fmt.Printf("%v\n", firstUser)

	secondUser, err := User{}.Find(2)
	if err != nil {
		log.Fatal("User.Find.Error", err)
	}
	fmt.Printf("%v\n", secondUser)

	Hanako, err := User{}.FindBy("Name", "Hanako")
	if err != nil {
		log.Fatal("User.FindBy.Error", err)
	}
	fmt.Printf("%v\n", Hanako)

	lastUser, _ := User{}.Last()
	err = lastUser.Delete()
	if err != nil {
		log.Fatal("user.Delete.Error", err)
	}

	kurapika, err := User{}.Create(UserParams{Name: "Kurapika", Age: 16})
	if err != nil {
		log.Fatal("User.Create.Error", err)
	}
	fmt.Printf("%v\n", kurapika)
}

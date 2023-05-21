# ayaorm

AYAORM is a golang ORM with an active record-like interface.
Instead of Ruby's metaprogramming, Ayaorm performs automatic code generation.

## How To Install

`go get -u github.com/okazaki-kk/ayaorm`

## Quick Start

Define a model struct in `user.go`

```go
package main

import "github.com/okazaki-kk/ayaorm"

// +AYAORM
type User struct {
 ayaorm.Schema
 Name string
 Age  int
}
```

Then, type `ayaorm user.go` and this will create new generated file `user_gen.go`

Finally, you can use ayaorm like this!

```go
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

 kurapika, err := User{}.Create(UserParams{Name: "Kurapika", Age: 16})
 if err != nil {
  log.Fatal("User.Create.Error", err)
 }
 fmt.Println(kurapika)

 kilua, err := User{}.Create(UserParams{Name: "Kilua", Age: 13})
 if err != nil {
  log.Fatal("User.Create.Error", err)
 }
 fmt.Println(kilua)

 fmt.Println("USER COUNT:", User{}.Count())

 users, err := User{}.All().Query()
 if err != nil {
  log.Fatal("User.All.Error", err)
 }
 for _, user := range users {
  fmt.Println(user)
 }

 users, err = User{}.Limit(2).Query()
 if err != nil {
  log.Fatal("User.All.Error", err)
 }
 for _, user := range users {
  fmt.Println(user)
 }

 reverseUsers, err := User{}.Order("Age", "desc").Query()
 if err != nil {
  log.Fatal("User.All.Error", err)
 }
 for _, user := range reverseUsers {
  fmt.Println(user)
 }

 Hanakos, err := User{}.Where("Name", "Hanako").Query()
 if err != nil {
  log.Fatal("User.All.Error", err)
 }
 for _, hanako := range Hanakos {
  fmt.Println(hanako)
 }

 newUser := User{Name: "Hanako", Age: 34}
 err = newUser.Save()
 if err != nil {
  log.Fatal("User.Save.Error", err)
 }
 fmt.Println(newUser)

 firstUser, err := User{}.First()
 if err != nil {
  log.Fatal("User.First.Error", err)
 }
 fmt.Println(firstUser)

 secondUser, err := User{}.Find(2)
 if err != nil {
  log.Fatal("User.Find.Error", err)
 }
 fmt.Println(secondUser)

 Hanako, err := User{}.FindBy("Name", "Hanako")
 if err != nil {
  log.Fatal("User.FindBy.Error", err)
 }
 fmt.Println(Hanako)

 lastUser, _ := User{}.Last()
 err = lastUser.Delete()
 if err != nil {
  log.Fatal("user.Delete.Error", err)
 }

 err = kurapika.Update(UserParams{Age: 18})
 if err != nil {
  log.Fatal("User.Update.Error", err)
 }
 fmt.Println(kurapika)
}
```

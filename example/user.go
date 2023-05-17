package main

// +AYAORM
type User struct {
	Id   int `db:"pk"`
	Name string
	Age  int
}

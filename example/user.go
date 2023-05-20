package main

import "github.com/okazaki-kk/ayaorm"

// +AYAORM
type User struct {
	ayaorm.Schema
	Name string
	Age  int
}

package main

import "github.com/okazaki-kk/ayaorm"

type Comment struct {
	ayaorm.Schema
	Content string
	Author  string
	PostId  int
}

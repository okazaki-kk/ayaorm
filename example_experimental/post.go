package main

import "github.com/okazaki-kk/ayaorm"

type Post struct {
	ayaorm.Schema
	Content string
	Author  string
}

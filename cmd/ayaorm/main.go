package main

import (
	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/template"
)

func main() {
	modelName, field := ayaorm.Inspect("./example/user.go")
	template.Generate(modelName, field)
}

package main

import (
	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/template"
)

func main() {
	modelName, fieldKeys, _ := ayaorm.Inspect("./example/user.go")
	filePath := "./main_gen.go"
	template.Generate(filePath, modelName, fieldKeys)
}

package main

import (
	"log"
	"os"
	"strings"

	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/template"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("please specify generate file")
	}
	from := os.Args[1]
	filePath := strings.Split(from, ".go")[0] + "_gen.go"

	modelName, fieldKeys, _ := ayaorm.Inspect(from)
	template.Generate(filePath, modelName, fieldKeys)
}

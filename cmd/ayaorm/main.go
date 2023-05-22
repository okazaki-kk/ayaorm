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

	fileInspect := ayaorm.Inspect(from)
	template.Generate(filePath, fileInspect[0].ModelName, fileInspect[0].FieldKeys)
}

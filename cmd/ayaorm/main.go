package main

import (
	"log"
	"os"

	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/template"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("please specify generate file")
	}
	from := os.Args[1]

	fileInspect := ayaorm.Inspect(from)
	template.Generate(fileInspect[0].ModelName, fileInspect[0].FieldKeys)
}

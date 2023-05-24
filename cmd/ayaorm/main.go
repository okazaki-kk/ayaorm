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
	for _, f := range fileInspect.StructInspect {
		if err := template.Generate(fileInspect.PackageName, f.ModelName, f.FieldKeys); err != nil {
			log.Fatal(err)
		}
	}

	err := template.GenerateDB()
	if err != nil {
		log.Fatal(err)
	}
}

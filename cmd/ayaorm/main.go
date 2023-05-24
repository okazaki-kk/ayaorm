package main

import (
	"log"
	"os"

	"github.com/okazaki-kk/ayaorm/template"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("please specify generate file")
	}
	from := os.Args[1]

	fileInspect := template.Inspect(from)
	err := template.Generate(fileInspect)
	if err != nil {
		log.Fatal(err)
	}

	err = template.GenerateDB()
	if err != nil {
		log.Fatal(err)
	}
}

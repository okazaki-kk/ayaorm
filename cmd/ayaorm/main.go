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

	fileInspect, err := template.Inspect(from)
	if err != nil {
		log.Fatal(err)
	}

	err = template.Generate(from, fileInspect)
	if err != nil {
		log.Fatal(err)
	}
}

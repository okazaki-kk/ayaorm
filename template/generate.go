package template

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

func Generate(filePath string, modelName string, field []string) {
	funcMap := template.FuncMap{
		"toSnakeCase": toSnakeCase,
	}

	var columns []string
	var columnNames []string
	for _, f := range field {
		columns = append(columns, f)
		columnNames = append(columnNames, toSnakeCase(f))
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(textBody)

	params := make(map[string]interface{})
	params["modelName"] = modelName
	params["snakeCaseModelName"] = toSnakeCase(modelName) + "s"
	params["columns"] = columns
	params["columnNames"] = columnNames

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("file create error: ", err)
	}
	defer file.Close()
	err = t.Execute(file, params)
	if err != nil {
		log.Fatal("template error: ", err)
	}

	err = exec.Command("go", "fmt", filePath).Run()
	if err != nil {
		log.Fatal("go fmt error: ", err)
	}
}

func toSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake))
}

package template

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

func Generate(modelName string, field map[string]string) {
	funcMap := template.FuncMap{
		"toSnakeCase": toSnakeCase,
	}

	var columns []string
	var columnNames []string
	for f := range field {
		columns = append(columns, f)
		columnNames = append(columnNames, toSnakeCase(f))
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(textBody)
	f, _ := os.Create("./main_gen.go")
	defer f.Close()

	params := make(map[string]interface{})
	params["modelName"] = modelName
	params["snakeCaseModelName"] = toSnakeCase(modelName) + "s"
	params["columns"] = columns
	params["columnNames"] = columnNames

	err := t.Execute(f, params)
	if err != nil {
		log.Fatal("template error: ", err)
	}

	err = exec.Command("go", "fmt", "./main_gen.go").Run()
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

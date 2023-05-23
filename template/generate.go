package template

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

func Generate(modelName string, field []string) error {
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
	t.New("Package").Parse(packageTextBody)
	t.New("Import").Parse(importTextBody)
	t.New("Relation").Parse(relationTextBody)
	t.New("Columns").Parse(columnsTextBody)

	params := make(map[string]interface{})
	params["modelName"] = modelName
	params["snakeCaseModelName"] = toSnakeCase(modelName) + "s"
	params["columns"] = columns
	params["columnNames"] = columnNames

	filePath := strings.ToLower(modelName) + "_gen.go"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = t.Execute(file, params)
	if err != nil {
		return err
	}

	err = exec.Command("go", "fmt", filePath).Run()
	if err != nil {
		return err
	}
	return nil
}

func GenerateDB() error {
	f, err := os.Create("db_gen.go")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(dbTextBody))
	if err != nil {
		return err
	}
	err = exec.Command("go", "fmt", "db_gen.go").Run()
	if err != nil {
		return err
	}
	return nil
}

func toSnakeCase(s string) string {
	const snake = "${1}_${2}"
	reg1 := regexp.MustCompile("([A-Z]+)([A-Z][a-z])")
	reg2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(reg2.ReplaceAllString(reg1.ReplaceAllString(s, snake), snake))
}

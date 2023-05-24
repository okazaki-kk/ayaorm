package template

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/okazaki-kk/ayaorm/template/templates"
)

func Generate(fileInspect FileInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": toSnakeCase,
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(TextBody)
	t.New("Package").Parse(templates.PackageTextBody)
	t.New("Import").Parse(templates.ImportTextBody)
	t.New("Relation").Parse(templates.RelationTextBody)
	t.New("Columns").Parse(templates.ColumnsTextBody)
	t.New("CRUD").Parse(templates.CrudTextBody)
	t.New("Search").Parse(templates.SearchTextBody)
	t.New("Query").Parse(templates.QueryTextBody)

	filePath := strings.ToLower(fileInspect.StructInspect[0].ModelName) + "_gen.go"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = t.Execute(file, fileInspect)
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

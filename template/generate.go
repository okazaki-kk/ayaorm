package template

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/okazaki-kk/ayaorm"
	"github.com/okazaki-kk/ayaorm/template/templates"
)

func Generate(from string, fileInspect FileInspect) error {
	filePath := strings.Split(from, ".go")[0] + "_gen.go"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("// Code generated by ayaorm. DO NOT EDIT.\npackage %s\n\n", fileInspect.PackageName))
	if err != nil {
		return err
	}

	var importText = `
	import (
		"fmt"

		"github.com/okazaki-kk/ayaorm"
	)
	`

	_, err = file.WriteString(importText)
	if err != nil {
		return err
	}

	for _, s := range fileInspect.StructInspect {
		s.PackageName = fileInspect.PackageName
		err := generateStruct(file, s)
		if err != nil {
			return err
		}
	}

	for _, f := range fileInspect.FuncInspect {
		var err error
		if f.BelongTo {
			err = generateBelongsToFunc(file, f)
		} else if f.HasMany {
			err = generateHasManyFunc(file, f)
		}
		if err != nil {
			return err
		}
	}

	params := generateValidateParams(fileInspect)
	for _, validate := range params {
		err := generateValidateFunc(file, validate)
		if err != nil {
			return err
		}
	}

	err = exec.Command("go", "fmt", filePath).Run()
	if err != nil {
		return err
	}

	err = generateDB(fileInspect.PackageName)
	if err != nil {
		return err
	}

	return nil
}

type templ struct {
	templateName string
	templateBody string
}

func generateStruct(file *os.File, structInspect StructInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": ayaorm.ToSnakeCase,
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(TextBody)

	var templs []templ = []templ{
		{
			templateName: "Relation",
			templateBody: templates.RelationTextBody,
		},
		{
			templateName: "Columns",
			templateBody: templates.ColumnsTextBody,
		},
		{
			templateName: "CRUD",
			templateBody: templates.CrudTextBody,
		},
		{
			templateName: "Search",
			templateBody: templates.SearchTextBody,
		},
		{
			templateName: "Query",
			templateBody: templates.QueryTextBody,
		},
	}

	for _, templ := range templs {
		_, err := t.New(templ.templateName).Parse(templ.templateBody)
		if err != nil {
			return err
		}
	}

	err := t.Execute(file, structInspect)
	if err != nil {
		return err
	}
	return nil
}

func generateHasManyFunc(file *os.File, funcInspect FuncInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": ayaorm.ToSnakeCase,
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(FuncBody)

	_, err := t.New("Joins").Parse(templates.HasManyTextBody)
	if err != nil {
		return err
	}

	err = t.Execute(file, funcInspect)
	if err != nil {
		return err
	}
	return nil
}

func generateBelongsToFunc(file *os.File, funcInspect FuncInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": ayaorm.ToSnakeCase,
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(FuncBody)

	_, err := t.New("Joins").Parse(templates.BelongsTextBody)
	if err != nil {
		return err
	}

	err = t.Execute(file, funcInspect)
	if err != nil {
		return err
	}
	return nil
}

type validate struct {
	Recv     string
	Name     string
	FuncName string
}

func generateValidateParams(fileInspect FileInspect) map[string][]validate {
	params := map[string][]validate{}

	for _, f := range fileInspect.FuncInspect {
		if f.ValidateLength {
			params[f.Recv] = append(params[f.Recv], validate{
				Recv:     f.Recv,
				Name:     f.ValidateLengthField(),
				FuncName: f.FuncName,
			})
		}
		if f.ValidatePresence {
			params[f.Recv] = append(params[f.Recv], validate{
				Recv:     f.Recv,
				Name:     f.ValidatePresenceField(),
				FuncName: f.FuncName,
			})
		}
	}
	return params
}

func generateValidateFunc(file *os.File, validates []validate) error {
	funcMap := template.FuncMap{
		"toSnakeCase": ayaorm.ToSnakeCase,
	}

	t, err := template.New("Base").Funcs(funcMap).Parse(FuncBody)
	if err != nil {
		return err
	}

	_, err = t.New("Joins").Parse(templates.ValidatePresenceTextBody)
	if err != nil {
		return err
	}

	err = t.Execute(file, validates)
	if err != nil {
		return err
	}
	return nil
}

func generateDB(packageName string) error {
	f, err := os.Create("db_gen.go")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("// Code generated by ayaorm. DO NOT EDIT.\npackage %s\n\n", packageName))
	if err != nil {
		return err
	}

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

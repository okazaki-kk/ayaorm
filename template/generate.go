package template

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	"github.com/okazaki-kk/ayaorm/template/templates"
	"github.com/okazaki-kk/ayaorm/utils"
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
		"errors"
		"fmt"

		"github.com/okazaki-kk/ayaorm"
		"github.com/okazaki-kk/ayaorm/utils"
		"github.com/okazaki-kk/ayaorm/validate"
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

	for _, f := range fileInspect.RelationFuncInspect {
		var err error
		if f.BelongTo {
			err = generateBelongsToFunc(file, f)
		} else if f.HasMany {
			err = generateHasManyFunc(file, f)
		} else if f.HasOne {
			err = generateHasOneFunc(file, f)
		}
		if err != nil {
			return err
		}
	}

	params := generateValidateParams(fileInspect)

	// sort by key and fix generate func order
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		err := generateValidateFunc(file, params[k], fileInspect.CustomRecv)
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

func generateStruct(file *os.File, structInspect StructInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": utils.ToSnakeCase,
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(TextBody)

	templs := []struct {
		templateName string
		templateBody string
	}{
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

func generateHasManyFunc(file *os.File, funcInspect RelationFuncInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": utils.ToSnakeCase,
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

func generateBelongsToFunc(file *os.File, funcInspect RelationFuncInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": utils.ToSnakeCase,
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

func generateHasOneFunc(file *os.File, funcInspect RelationFuncInspect) error {
	funcMap := template.FuncMap{
		"toSnakeCase": utils.ToSnakeCase,
	}

	t, _ := template.New("Base").Funcs(funcMap).Parse(FuncBody)

	_, err := t.New("Joins").Parse(templates.HasOneTextBody)
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
	Name     string
	FuncName string
}

type validates struct {
	Validates        []validate
	CustomValidation bool
	Recv             string
}

func generateValidateParams(fileInspect FileInspect) map[string]validates {
	params := map[string]validates{}

	for _, f := range fileInspect.StructInspect {
		params[f.ModelName] = validates{Validates: []validate{{}}, Recv: f.ModelName}
	}

	for _, f := range fileInspect.ValidateFuncInspect {
		if f.ValidateLength {
			v := params[f.Recv]
			v.Validates = append(v.Validates, validate{Name: f.ValidateLengthField(), FuncName: f.FuncName})
			params[f.Recv] = v
		}
		if f.ValidatePresence {
			v := params[f.Recv]
			v.Validates = append(v.Validates, validate{Name: f.ValidatePresenceField(), FuncName: f.FuncName})
			params[f.Recv] = v
		}
		if f.ValidateNumericality {
			v := params[f.Recv]
			v.Validates = append(v.Validates, validate{Name: f.ValidateNumericalityField(), FuncName: f.FuncName})
			params[f.Recv] = v
		}
	}
	return params
}

func generateValidateFunc(file *os.File, validates validates, customRecv []string) error {
	funcMap := template.FuncMap{
		"toSnakeCase": utils.ToSnakeCase,
	}
	validates.CustomValidation = utils.Contains(customRecv, validates.Recv)

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

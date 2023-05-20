package ayaorm

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	filePath := "./temp_inspect.go"
	file, err := os.Create(filePath)
	defer os.Remove(filePath)
	if err != nil {
		log.Fatal("file create error: ", err)
	}
	defer file.Close()

	var userStruct = `package main

	// +AYAORM
	type User struct {
		Id   int ` + "`" + `db:"pk"` + "`" + `
		Name string
		Age  int
	}
	`

	_, err = file.Write([]byte(userStruct))
	assert.NoError(t, err)

	modelName, fieldKeys, fieldValues := Inspect(filePath)
	assert.Equal(t, "User", modelName)
	assert.Equal(t, []string{"Id", "Name", "Age"}, fieldKeys)
	assert.Equal(t, []string{"int", "string", "int"}, fieldValues)
}

package template

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
	type Post struct {
		ayaorm.Schema
		Content string
		Author  string
	}

	type Comment struct {
		ayaorm.Schema
		Content string
		Author  string
		PostId  int
	}

	func (m Post) hasManyPosts() {
	}
	`

	_, err = file.Write([]byte(userStruct))
	assert.NoError(t, err)

	fileInspect := Inspect(filePath)

	assert.Equal(t, "main", fileInspect.PackageName)
	assert.Equal(t, 2, len(fileInspect.StructInspect))
	assert.Equal(t, 1, len(fileInspect.FuncInspect))

	assert.Equal(t, "Post", fileInspect.StructInspect[0].ModelName)
	assert.Equal(t, []string{"Id", "Content", "Author", "CreatedAt", "UpdatedAt"}, fileInspect.StructInspect[0].FieldKeys)
	assert.Equal(t, []string{"int", "string", "string", "time.Time", "time.Time"}, fileInspect.StructInspect[0].FieldValues)

	assert.Equal(t, "Comment", fileInspect.StructInspect[1].ModelName)
	assert.Equal(t, []string{"Id", "Content", "Author", "PostId", "CreatedAt", "UpdatedAt"}, fileInspect.StructInspect[1].FieldKeys)
	assert.Equal(t, []string{"int", "string", "string", "int", "time.Time", "time.Time"}, fileInspect.StructInspect[1].FieldValues)

	assert.Equal(t, "hasManyPosts", fileInspect.FuncInspect[0].FuncName)
	assert.Equal(t, "Post", fileInspect.FuncInspect[0].Recv)
}
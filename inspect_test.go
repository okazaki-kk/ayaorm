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

	packageName, fileInspect, _ := Inspect(filePath)
	assert.Equal(t, "main", packageName)
	assert.Equal(t, "Post", fileInspect[0].ModelName)
	assert.Equal(t, []string{"Id", "Content", "Author", "CreatedAt", "UpdatedAt"}, fileInspect[0].FieldKeys)
	assert.Equal(t, []string{"int", "string", "string", "time.Time", "time.Time"}, fileInspect[0].FieldValues)

	assert.Equal(t, "Comment", fileInspect[1].ModelName)
	assert.Equal(t, []string{"Id", "Content", "Author", "PostId", "CreatedAt", "UpdatedAt"}, fileInspect[1].FieldKeys)
	assert.Equal(t, []string{"int", "string", "string", "int", "time.Time", "time.Time"}, fileInspect[1].FieldValues)

	var funcStruct = `package ayaorm

	// +AYAORM
	func (m Post) hasManyComments() {
	}
	`

	filePath1 := "./temp_inspect.go"
	file, err = os.Create(filePath1)
	assert.NoError(t, err)
	defer os.Remove(filePath)
	_, err = file.Write([]byte(funcStruct))
	assert.NoError(t, err)

	packageName, _, funcInspect := Inspect(filePath)
	assert.Equal(t, "ayaorm", packageName)
	assert.Equal(t, "hasManyComments", funcInspect.FuncName)
	assert.Equal(t, "Post", funcInspect.Recv)
}

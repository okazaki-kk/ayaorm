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

	var userStruct = `package testss

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

	func (m Post) hasManyComments() {
	}

	func (m Comment) belongsToPost() {
	}

	func (m Post) validatesPresenceOfAuthor() ayaorm.Rule {
		return ayaorm.MakeRule().Presence()
	}

	func (m Post) validateLengthOfContent() ayaorm.Rule {
		return ayaorm.MakeRule().MaxLength(10).MinLength(3)
	}

	func (m User) validateNumericalityOfAge() ayaorm.Rule {
		return ayaorm.MakeRule().Numericality().OnlyInteger()
	}

	func (m User) validateCustomRule() ayaorm.Rule {
		return ayaorm.CustomRule(func(es *[]error) {
			if m.Name == "custom-example" {
				*es = append(*es, errors.New("name must not be custom-example"))
			}
		})
	}
	`

	_, err = file.Write([]byte(userStruct))
	assert.NoError(t, err)

	fileInspect := Inspect(filePath)

	assert.Equal(t, "testss", fileInspect.PackageName)
	assert.Equal(t, 2, len(fileInspect.StructInspect))
	assert.Equal(t, 5, len(fileInspect.FuncInspect))
	assert.Equal(t, 1, len(fileInspect.CustomRecv))
	assert.Equal(t, "User", fileInspect.CustomRecv[0])

	assert.Equal(t, "Post", fileInspect.StructInspect[0].ModelName)
	assert.Equal(t, []string{"Id", "Content", "Author", "CreatedAt", "UpdatedAt"}, fileInspect.StructInspect[0].FieldKeys)
	assert.Equal(t, []string{"int", "string", "string", "time.Time", "time.Time"}, fileInspect.StructInspect[0].FieldValues)

	assert.Equal(t, "Comment", fileInspect.StructInspect[1].ModelName)
	assert.Equal(t, []string{"Id", "Content", "Author", "PostId", "CreatedAt", "UpdatedAt"}, fileInspect.StructInspect[1].FieldKeys)
	assert.Equal(t, []string{"int", "string", "string", "int", "time.Time", "time.Time"}, fileInspect.StructInspect[1].FieldValues)

	assert.Equal(t, "hasManyComments", fileInspect.FuncInspect[0].FuncName)
	assert.Equal(t, "Post", fileInspect.FuncInspect[0].Recv)
	assert.Equal(t, true, fileInspect.FuncInspect[0].HasMany)
	assert.Equal(t, "Comment", fileInspect.FuncInspect[0].HasManyModel())

	assert.Equal(t, "belongsToPost", fileInspect.FuncInspect[1].FuncName)
	assert.Equal(t, "Comment", fileInspect.FuncInspect[1].Recv)
	assert.Equal(t, true, fileInspect.FuncInspect[1].BelongTo)
	assert.Equal(t, "Post", fileInspect.FuncInspect[1].BelongsToModel())

	assert.Equal(t, "validatesPresenceOfAuthor", fileInspect.FuncInspect[2].FuncName)
	assert.Equal(t, "Post", fileInspect.FuncInspect[2].Recv)
	assert.Equal(t, true, fileInspect.FuncInspect[2].ValidatePresence)
	assert.Equal(t, "Author", fileInspect.FuncInspect[2].ValidatePresenceField())

	assert.Equal(t, "validateLengthOfContent", fileInspect.FuncInspect[3].FuncName)
	assert.Equal(t, "Post", fileInspect.FuncInspect[3].Recv)
	assert.Equal(t, true, fileInspect.FuncInspect[3].ValidateLength)
	assert.Equal(t, "Content", fileInspect.FuncInspect[3].ValidateLengthField())

	assert.Equal(t, "validateNumericalityOfAge", fileInspect.FuncInspect[4].FuncName)
	assert.Equal(t, "User", fileInspect.FuncInspect[4].Recv)
	assert.Equal(t, true, fileInspect.FuncInspect[4].ValidateNumericality)
	assert.Equal(t, "Age", fileInspect.FuncInspect[4].ValidateNumericalityField())
}

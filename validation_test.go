package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestPost struct {
	Schema
	Author  string
	Content string
}

func (m TestPost) validatesPresenceOfAuthor() Rule {
	return MakeRule().Presence()
}

func (m TestPost) validateMaxLengthOfContent() Rule {
	return MakeRule().MaxLength(10)
}

func TestIsValid(t *testing.T) {
	validator := NewValidator(TestPost{}.validatesPresenceOfAuthor().Rule())

	t.Run("when value is empty", func(t *testing.T) {
		result, errors := validator.IsValid("Author", "")
		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Author can't be blank", errors[0].Error())
	})

	t.Run("when value is not empty", func(t *testing.T) {
		result, errors := validator.IsValid("Author", "okazaki")
		assert.Equal(t, true, result)
		assert.Equal(t, 0, len(errors))
	})

	validator = NewValidator(TestPost{}.validateMaxLengthOfContent().Rule())
	t.Run("when text is too long", func(t *testing.T) {
		result, errors := validator.IsValid("Content", "abcdefghijklmn")

		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Content is too long (maximum is 10 characters)", errors[0].Error())
	})

	t.Run("when text is not too long", func(t *testing.T) {
		result, errors := validator.IsValid("Content", "abcde")

		assert.Equal(t, true, result)
		assert.Equal(t, 0, len(errors))
	})
}

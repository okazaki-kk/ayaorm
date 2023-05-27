package ayaorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestPost struct {
	Schema
	Author  string
	Content string
	Age     int
}

func (m TestPost) validatesPresenceOfAuthor() Rule {
	return MakeRule().Presence()
}

func (m TestPost) validateLengthOfContent() Rule {
	return MakeRule().MaxLength(10).MinLength(3)
}

func (m TestPost) validateNumericalityOfAge() Rule {
	return MakeRule().Numericality().OnlyInteger()
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

	validator = NewValidator(TestPost{}.validateLengthOfContent().Rule())
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

	t.Run("when text is too short", func(t *testing.T) {
		result, errors := validator.IsValid("Content", "ab")

		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Content is too short (minimum is 3 characters)", errors[0].Error())
	})

	validator = NewValidator(TestPost{}.validateNumericalityOfAge().Rule())
	t.Run("when age is not numerical", func(t *testing.T) {
		result, errors := validator.IsValid("Age", "20.0")

		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Age must be number", errors[0].Error())
	})

	t.Run("when age is not integer", func(t *testing.T) {
		result, errors := validator.IsValid("Age", 20.0)

		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Age must be integer", errors[0].Error())
	})

	t.Run("when age is numerical", func(t *testing.T) {
		result, errors := validator.IsValid("Age", 20)

		assert.Equal(t, true, result)
		assert.Equal(t, 0, len(errors))
	})
}

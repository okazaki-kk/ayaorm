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
}

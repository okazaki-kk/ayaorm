package validate

import (
	"testing"

	"github.com/okazaki-kk/ayaorm"
	"github.com/stretchr/testify/assert"
)

type TestPost struct {
	ayaorm.Schema
	Author  string
	Content string
	Age     int
	Age1    int
}

func (m TestPost) validatesPresenceOfAuthor() Rule {
	return MakeRule().Presence()
}

func (m TestPost) validateLengthOfContent() Rule {
	return MakeRule().MaxLength(10).MinLength(3)
}

func (m TestPost) validateNumericalityOfAge() Rule {
	return MakeRule().Numericality().Positive()
}

func (m TestPost) validateNumericalityOfAge1() Rule {
	return MakeRule().Numericality().Negative()
}

func (m TestPost) validateNumericalityOfAge2() Rule {
	return MakeRule().Numericality().Positive().OnCreate()
}

func (m TestPost) validateNumericalityOfAge3() Rule {
	return MakeRule().Numericality().Positive().OnUpdate()
}

func TestIsValid(t *testing.T) {
	validator := NewValidator(TestPost{}.validatesPresenceOfAuthor().Rule())
	assert.True(t, validator.rule.onCreate)
	assert.True(t, validator.rule.onUpdate)

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
	assert.True(t, validator.rule.onCreate)
	assert.True(t, validator.rule.onUpdate)
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
	assert.True(t, validator.rule.onCreate)
	assert.True(t, validator.rule.onUpdate)
	t.Run("when age is not numerical", func(t *testing.T) {
		result, errors := validator.IsValid("Age", "20.0")

		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Age must be number", errors[0].Error())
	})

	t.Run("when age is positive", func(t *testing.T) {
		result, errors := validator.IsValid("Age", 20)

		assert.Equal(t, true, result)
		assert.Equal(t, 0, len(errors))
	})

	t.Run("when age is negative", func(t *testing.T) {
		result, errors := validator.IsValid("Age", -20)

		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Age must be positive", errors[0].Error())
	})

	validator = NewValidator(TestPost{}.validateNumericalityOfAge1().Rule())
	assert.True(t, validator.rule.onCreate)
	assert.True(t, validator.rule.onUpdate)

	t.Run("when age1 is negative", func(t *testing.T) {
		result, errors := validator.IsValid("Age1", -20)

		assert.Equal(t, true, result)
		assert.Equal(t, 0, len(errors))
	})

	t.Run("when age1 is positive", func(t *testing.T) {
		result, errors := validator.IsValid("Age1", 20)

		assert.Equal(t, false, result)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "Age1 must be negative", errors[0].Error())
	})

	validator = NewValidator(TestPost{}.validateNumericalityOfAge2().Rule())
	assert.True(t, validator.rule.onCreate)
	assert.False(t, validator.rule.onUpdate)

	validator = NewValidator(TestPost{}.validateNumericalityOfAge3().Rule())
	assert.False(t, validator.rule.onCreate)
	assert.True(t, validator.rule.onUpdate)
}

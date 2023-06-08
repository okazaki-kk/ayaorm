package validate

import (
	"fmt"
)

type Rule interface {
	Rule() *Validation
}

func MakeRule() *Validation {
	return &Validation{}
}

type Validation struct {
	custom       customValidation
	presence     *presence
	maxLength    *maxLength
	minLength    *minLength
	numericality *numericality
	positive     *positive
	negative     *negative
	onCreate     bool
	onUpdate     bool
}

func (v *Validation) Rule() *Validation {
	return v
}

func NewValidator(rule *Validation) Validator {
	if !rule.onCreate && !rule.onUpdate {
		rule.onCreate = true
		rule.onUpdate = true
	}
	return Validator{rule}
}

type Validator struct {
	rule *Validation
}

func (v Validator) IsValid(name string, value interface{}) (bool, []error) {
	if v.rule == nil {
		return true, nil
	}

	result := true
	errors := []error{}

	if v.rule.presence != nil && v.rule.presence.presence {
		if ok, err := v.isPresent(name, value); !ok {
			result = false
			errors = append(errors, err)
		}
	}

	if v.rule.maxLength != nil && len(value.(string)) > v.rule.maxLength.maxLength {
		result = false
		errors = append(errors, fmt.Errorf("%s is too long (maximum is %d characters)", name, v.rule.maxLength.maxLength))
	}
	if v.rule.minLength != nil && len(value.(string)) < v.rule.minLength.minLength {
		result = false
		errors = append(errors, fmt.Errorf("%s is too short (minimum is %d characters)", name, v.rule.minLength.minLength))
	}

	if v.rule.numericality != nil && v.rule.numericality.numericality {
		if ok, err := v.isNumericality(name, value); !ok {
			result = false
			errors = append(errors, err)
		}
	}

	return result, errors
}

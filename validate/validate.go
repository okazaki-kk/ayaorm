package validate

import (
	"fmt"

	"github.com/okazaki-kk/ayaorm"
)

type Rule interface {
	Rule() *Validation
}

func MakeRule() *Validation {
	return &Validation{}
}

func CustomRule(c CustomValidation) *Validation {
	v := &Validation{}
	return v.CustomRule(c)
}

func (v *Validation) CustomRule(c CustomValidation) *Validation {
	v.custom = c
	return v
}

type CustomValidation func(errors *[]error)

type Validation struct {
	custom       CustomValidation
	presence     *presence
	maxLength    *maxLength
	minLength    *minLength
	numericality *numericality
	positive     *positive
	negative     *negative
}

func (v *Validation) Rule() *Validation {
	return v
}

func (v *Validation) Presence() *presence {
	if v.presence == nil {
		v.presence = newPresence(v)
	}
	return v.presence
}

type presence struct {
	*Validation
	presence bool
}

func newPresence(v *Validation) *presence {
	return &presence{
		Validation: v,
		presence:   true,
	}
}

func (p *presence) Rule() *Validation {
	return p.Validation
}

func (v *Validation) MaxLength(max int) *maxLength {
	if v.maxLength == nil {
		v.maxLength = newMaxLength(v, max)
	}
	return v.maxLength
}

type maxLength struct {
	*Validation
	maxLength int
}

func newMaxLength(v *Validation, max int) *maxLength {
	return &maxLength{
		Validation: v,
		maxLength:  max,
	}
}

func (l *maxLength) Rule() *Validation {
	return l.Validation
}

func (v *Validation) MinLength(min int) *minLength {
	if v.minLength == nil {
		v.minLength = newMinLength(v, min)
	}
	return v.minLength
}

type minLength struct {
	*Validation
	minLength int
}

func newMinLength(v *Validation, min int) *minLength {
	return &minLength{
		Validation: v,
		minLength:  min,
	}
}

func (l *minLength) Rule() *Validation {
	return l.Validation
}

func (v *Validation) Numericality() *numericality {
	if v.numericality == nil {
		v.numericality = newNumericality(v)
	}
	return v.numericality
}

type numericality struct {
	*Validation
	numericality bool
}

func newNumericality(v *Validation) *numericality {
	return &numericality{
		Validation:   v,
		numericality: true,
	}
}

func (n *numericality) Rule() *Validation {
	return n.Validation
}

func (v *Validation) Positive() *positive {
	if v.positive == nil {
		v.positive = newPositive(v)
	}
	return v.positive
}

type positive struct {
	*Validation
	positive bool
}

func newPositive(v *Validation) *positive {
	return &positive{
		Validation: v,
		positive:   true,
	}
}

func (n *positive) Rule() *Validation {
	return n.Validation
}

func (v *Validation) Negative() *negative {
	if v.negative == nil {
		v.negative = newNegative(v)
	}
	return v.negative
}

type negative struct {
	*Validation
	negative bool
}

func newNegative(v *Validation) *negative {
	return &negative{
		Validation: v,
		negative:   true,
	}
}

func (n *negative) Rule() *Validation {
	return n.Validation
}

func NewValidator(rule *Validation) Validator {
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

func (v Validator) isNumericality(name string, value interface{}) (bool, error) {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		if v.rule.positive != nil && v.rule.positive.positive {
			// TODO: type assertion for various types
			if value.(int) < 0 {
				return false, fmt.Errorf("%s must be positive", name)
			}
		}
		if v.rule.negative != nil && v.rule.negative.negative {
			// TODO: type assertion for various types
			if value.(int) > 0 {
				return false, fmt.Errorf("%s must be negative", name)
			}
		}
		return true, nil
	default:
		return false, fmt.Errorf("%s must be number", name)
	}
}

func (v Validator) isPresent(name string, value interface{}) (bool, error) {
	if ayaorm.IsZero(value) {
		return false, fmt.Errorf("%s can't be blank", name)
	}
	return true, nil
}

func (v Validator) Custom() CustomValidation {
	if v.rule.custom == nil {
		return func(errors *[]error) {}
	}
	return v.rule.custom
}

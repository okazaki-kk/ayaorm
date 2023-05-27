package ayaorm

import (
	"fmt"
	"reflect"
)

type Rule interface {
	Rule() *Validation
}

func MakeRule() *Validation {
	return &Validation{}
}

type Validation struct {
	presence     *presence
	maxLength    *maxLength
	numericality *numericality
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

type Validator struct {
	rule *Validation
}

func NewValidator(rule *Validation) Validator {
	return Validator{rule}
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

	if v.rule.maxLength != nil {
		if reflect.TypeOf(value).Kind() != reflect.String {
			result = false
			errors = append(errors, fmt.Errorf("%s must be string", name))
		} else {
			if len(value.(string)) > v.rule.maxLength.maxLength {
				result = false
				errors = append(errors, fmt.Errorf("%s is too long (maximum is %d characters)", name, v.rule.maxLength.maxLength))
			}
		}
	}

	if v.rule.numericality != nil && v.rule.numericality.numericality {
		switch value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			return true, nil
		case float32, float64:
			return true, nil
		}
		return false, []error{fmt.Errorf("%s must be number", name)}
	}

	return result, errors
}

func (v Validator) isPresent(name string, value interface{}) (bool, error) {
	if IsZero(value) {
		return false, fmt.Errorf("%s can't be blank", name)
	}
	return true, nil
}

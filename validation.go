package ayaorm

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
	presence *presence
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

	if v.rule.presence.presence {
		if ok, err := v.isPresent(name, value); !ok {
			result = false
			errors = append(errors, err)
		}
	}
	return result, errors
}

func (v Validator) isPresent(name string, value interface{}) (bool, error) {
	if IsZero(value) {
		return false, fmt.Errorf("%s can't be blank", name)
	}
	return true, nil
}

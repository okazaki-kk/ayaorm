package validate

import "fmt"

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

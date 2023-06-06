package validate

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

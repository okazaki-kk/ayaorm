package validate

type customValidation func(errors *[]error)

func (v Validator) Custom() customValidation {
	if v.rule.custom == nil {
		return func(errors *[]error) {}
	}
	return v.rule.custom
}

func CustomRule(c customValidation) *Validation {
	v := &Validation{}
	return v.newCustomRule(c)
}

func (v *Validation) newCustomRule(c customValidation) *Validation {
	v.custom = c
	return v
}

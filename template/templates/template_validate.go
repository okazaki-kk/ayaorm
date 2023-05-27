package templates

var ValidatePresenceTextBody = `
	func (m Post) IsValid() (bool, []error) {
		result := true
		var errors []error

		rules := map[string]*ayaorm.Validation{
			"author": m.validatesPresenceOfAuthor().Rule(),
		}

		for name, rule := range rules {
			if ok, errs := ayaorm.NewValidator(rule).IsValid(name, m.fieldValuesByName(name)); !ok {
				result = false
				errors = append(errors, errs...)
			}
		}

		if len(errors) > 0 {
			result = false
		}
		return result, errors
	}
`
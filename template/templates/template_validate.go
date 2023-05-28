package templates

var ValidatePresenceTextBody = `
	func (m {{.Recv}}) IsValid() (bool, []error) {
		result := true
		var errors []error

		rules := map[string]*ayaorm.Validation{
			"{{toSnakeCase .ValidatePresenceField}}": m.{{.FuncName}}().Rule(),
			"content": m.validateLengthOfContent().Rule(),
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

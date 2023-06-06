package templates

var ValidatePresenceTextBody = `
	func (m {{.Recv}}) IsValid() (bool, []error) {
		result := true
		var errors []error

		rules := map[string]*validate.Validation{
			{{ range $key, $value := .Validates -}}
			{{ if eq $value.Name "" -}} {{continue}} {{ end -}}
			"{{toSnakeCase $value.Name}}": m.{{$value.FuncName}}().Rule(),
			{{ end -}}
		}

		for name, rule := range rules {
			if ok, errs := validate.NewValidator(rule).IsValid(name, m.fieldValuesByName(name)); !ok {
				result = false
				errors = append(errors, errs...)
			}
		}

		{{ if .CustomValidation }}
		customs := []*validate.Validation{m.validateCustomRule().Rule()}
		for _, rule := range customs {
			custom := validate.NewValidator(rule).Custom()
			custom(&errors)
		}
		{{ end }}

		if len(errors) > 0 {
			result = false
		}
		return result, errors
	}
`

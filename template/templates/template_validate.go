package templates

var ValidatePresenceTextBody = `
	{{ $struct := index . 0 }}
	func (m {{$struct.Recv}}) IsValid() (bool, []error) {
		result := true
		var errors []error

		rules := map[string]*ayaorm.Validation{
			{{ range $key, $value := . -}}
			{{ if eq $value.Name "" -}} {{continue}} {{ end -}}
			"{{toSnakeCase $value.Name}}": m.{{$value.FuncName}}().Rule(),
			{{ end -}}
		}

		for name, rule := range rules {
			if ok, errs := ayaorm.NewValidator(rule).IsValid(name, m.fieldValuesByName(name)); !ok {
				result = false
				errors = append(errors, errs...)
			}
		}

		{{ if .NeedCustom }}
		customs := []*ayaorm.Validation{m.validateCustomRule().Rule()}
		for _, rule := range customs {
			custom := ayaorm.NewValidator(rule).Custom()
			custom(&errors)
		}
		{{ end }}

		if len(errors) > 0 {
			result = false
		}
		return result, errors
	}
`

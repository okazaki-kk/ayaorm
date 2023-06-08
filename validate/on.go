package validate

type On struct {
	OnCreate bool
	OnUpdate bool
}

func (v Validator) On(on On) Validator {
	if v.rule != nil && !v.rule.onCreate && on.OnCreate {
		v.rule = nil
	}
	if v.rule != nil && !v.rule.onUpdate && on.OnUpdate {
		v.rule = nil
	}
	return v
}

func (v *Validation) OnCreate() *Validation {
	v.onCreate = true
	v.onUpdate = false
	return v
}

func (v *Validation) OnUpdate() *Validation {
	v.onUpdate = true
	v.onCreate = false
	return v
}

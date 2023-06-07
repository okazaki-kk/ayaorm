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

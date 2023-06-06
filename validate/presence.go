package validate

import (
	"fmt"

	"github.com/okazaki-kk/ayaorm"
)

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

func (v Validator) isPresent(name string, value interface{}) (bool, error) {
	if ayaorm.IsZero(value) {
		return false, fmt.Errorf("%s can't be blank", name)
	}
	return true, nil
}

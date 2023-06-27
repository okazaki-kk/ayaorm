package null

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type NullFloat struct {
	i sql.NullFloat64
}

func (i NullFloat) String() string {
	if !i.i.Valid {
		return ""
	}
	return fmt.Sprint(i.i.Float64)
}

func (i *NullFloat) Valid() bool {
	return i.i.Valid
}

func (i *NullFloat) Set(value float64) *NullFloat {
	i.i.Valid = true
	i.i.Float64 = value
	return i
}

func (i *NullFloat) Scan(value interface{}) error {
	return i.i.Scan(value)
}

func (i NullFloat) Value() (driver.Value, error) {
	if !i.Valid() {
		return nil, nil
	}
	return i.i.Float64, nil
}

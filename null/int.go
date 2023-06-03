package null

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Nullint struct {
	i sql.NullInt64
}

func (i Nullint) String() string {
	if !i.i.Valid {
		return ""
	}
	return fmt.Sprint(i.i.Int64)
}

func (i *Nullint) Valid() bool {
	return i.i.Valid
}

func (i *Nullint) Set(value int64) *Nullint {
	i.i.Valid = true
	i.i.Int64 = value
	return i
}

func (i *Nullint) Scan(value interface{}) error {
	return i.i.Scan(value)
}

func (i Nullint) Value() (driver.Value, error) {
	if !i.Valid() {
		return nil, nil
	}
	return i.i.Int64, nil
}

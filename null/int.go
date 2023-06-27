package null

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type NullInt struct {
	i sql.NullInt64
}

func (i NullInt) String() string {
	if !i.i.Valid {
		return ""
	}
	return fmt.Sprint(i.i.Int64)
}

func (i *NullInt) Valid() bool {
	return i.i.Valid
}

func (i *NullInt) Set(value int64) *NullInt {
	i.i.Valid = true
	i.i.Int64 = value
	return i
}

func (i *NullInt) Scan(value interface{}) error {
	return i.i.Scan(value)
}

func (i NullInt) Value() (driver.Value, error) {
	if !i.Valid() {
		return nil, nil
	}
	return i.i.Int64, nil
}

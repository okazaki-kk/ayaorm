package null

import "database/sql/driver"

type Null interface {
	Strint() string
	Valid() bool
	Set(string)
	Scan(interface{}) error
	Value() (driver.Value, error)
}

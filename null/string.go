package null

import (
	"database/sql"
	"database/sql/driver"
)

type NullString struct {
	s sql.NullString
}

func (s NullString) String() string {
	if !s.s.Valid {
		return ""
	}
	return s.s.String
}

func (s *NullString) Valid() bool {
	return s.s.Valid
}

func (s *NullString) Set(value string) {
	s.s.Valid = true
	s.s.String = value
}

func (s *NullString) Scan(value interface{}) error {
	return s.s.Scan(value)
}

func (s NullString) Value() (driver.Value, error) {
	if !s.Valid() {
		return nil, nil
	}
	return s.s.String, nil
}

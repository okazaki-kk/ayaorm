package ayaorm

import "time"

type Schema struct {
	Id        int `db:"pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package entities

import (
	"time"
)

// Employee зачем uri, db, binding ?
type Employee struct {
	Lastname  string     `uri:"lastname" db:"lastname" binding:"required"`
	StartDate *time.Time `db:"startdate"`
}

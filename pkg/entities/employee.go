package entities

import (
	"time"
)

type Employee struct {
	FIO       string     `uri:"fio" db:"lastname" binding:"required"`
	StartDate *time.Time `db:"startdate"`
}

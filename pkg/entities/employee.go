package entities

import (
	"time"
)

type Employee struct {
	FIO       string     `uri:"fio" db:"lastname" binding:"required"`
	StartDate *time.Time `db:"startdate"`
}

func (e *Employee) NewEmployee(name string, year, month, day int) {
	e.FIO = name
	e.StartDate = e.setDate(year, month, day)
}
func (e *Employee) setDate(year, month, day int) *time.Time {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &date
}

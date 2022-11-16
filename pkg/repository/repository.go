package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/m0n7h0ff/date-calc/pkg/entities"
)

type Authorization interface {
}

type Employee interface {
	GetAllEmployee() []entities.Employee
	GetByName(name string) (entities.Employee, error)
}

type Repository struct {
	Authorization
	Employee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Employee: NewEmployeePostgres(db),
	}
}

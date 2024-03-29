package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/m0n7h0ff/date-calc/pkg/entities"
)

type Device interface {
}

type Employee interface {
	GetAllEmployee() ([]entities.Employee, error)
	GetByLastName(lastname string) (entities.Employee, error)
}

type Repository struct {
	Employee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Employee: NewEmployeePostgres(db),
	}
}

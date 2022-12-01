package service

import (
	"github.com/m0n7h0ff/date-calc/pkg/entities"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
)

type Employee interface {
	GetAllEmployee() ([]entities.Employee, error)
	GetByLastName(lastname string) (entities.Employee, error)
}

type Service struct {
	Employee
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(repos.Employee),
	}
}

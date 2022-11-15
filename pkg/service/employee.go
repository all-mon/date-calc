package service

import (
	"github.com/m0n7h0ff/date-calc/pkg/entities"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
)

type EmployeeService struct {
	repo repository.Employee
}

func NewEmployeeService(repo repository.Employee) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetAllEmployee() []entities.Employee{
	return s.repo.GetAllEmployee()
}

func (s *EmployeeService) GetByName(name string) entities.Employee{
	return s.repo.GetByName(name)
}

package service

import (
	"github.com/m0n7h0ff/date-calc/pkg/entities"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
)

type Authorization interface {
}

type Employee interface {
	GetAllEmployee() []entities.Employee
	GetByName(name string) entities.Employee
}

type Service struct {
	Authorization
	Employee
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
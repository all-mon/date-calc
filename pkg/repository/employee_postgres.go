package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/m0n7h0ff/date-calc/pkg/entities"
)

type EmployeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}

func (r *EmployeePostgres) GetAllEmployee() ([]entities.Employee, error) {
	var employees []entities.Employee
	query := fmt.Sprintf("SELECT * FROM"+" %s", employeeTable)
	err := r.db.Select(employees, query)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeePostgres) GetByLastName(lastname string) (entities.Employee, error) {
	findEmployee := entities.Employee{}
	query := fmt.Sprintf("SELECT lastname, startdate FROM"+" %s WHERE lastname = $1 ", employeeTable)
	fmt.Println(query)
	err := r.db.Get(&findEmployee, query, lastname)
	if err != nil {
		fmt.Println(err.Error())
		return findEmployee, err
	}
	fmt.Println(findEmployee)
	return findEmployee, nil
}

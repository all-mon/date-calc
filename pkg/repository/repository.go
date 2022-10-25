package repository

import (
	date_calc "github.com/m0n7h0ff/date-calc/pkg/entities"
)

func GetEmployeeList() map[string]date_calc.Employee {
	emp1 := new(date_calc.Employee)
	emp1.NewEmployee("Монахов", 2022, 6, 3)

	emp2 := new(date_calc.Employee)
	emp2.NewEmployee("Дубинин", 2022, 6, 5)

	emp3 := new(date_calc.Employee)
	emp3.NewEmployee("Тюшняков", 2022, 6, 1)

	emp4 := new(date_calc.Employee)
	emp4.NewEmployee("Перехода", 2022, 6, 7)

	mapOfEmployee := make(map[string]date_calc.Employee)
	mapOfEmployee[emp1.FIO] = *emp1
	mapOfEmployee[emp2.FIO] = *emp2
	mapOfEmployee[emp3.FIO] = *emp3
	mapOfEmployee[emp4.FIO] = *emp4

	return mapOfEmployee
}

package main

import (
	"bufio"
	"fmt"
	date_calc "github.com/m0n7h0ff/date-calc"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	mapOfEmployee := setEmployeeList()
	ReadAndPrint(mapOfEmployee)
	//что-бы консоль не закрывалась сразу
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}

func ReadAndPrint(mapOfEmployee map[string]date_calc.Employee) {
	enteredName := readInput("Введите фамилию:\t")

	foundEmployee, ok := mapOfEmployee[enteredName]
	if !ok {
		log.Fatal("Имени нет")
	}
	startDate := *foundEmployee.StartDate
	enteredDate := readInput("Введите дату (формат 2022-06-30):\t")

	parseDate, err := time.Parse("2006-01-02", enteredDate)
	if err != nil {
		log.Printf("Неправильный формат даты: %v \n", err)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
	}

	differenceBetweenDates := getDifferenceBetweenDates(startDate, parseDate)
	_, fractional := math.Modf(float64(differenceBetweenDates) / 8.0) //остаток
	fmt.Println()
	printResult(fractional)
}

func setEmployeeList() map[string]date_calc.Employee {
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

func getDifferenceBetweenDates(startDate time.Time, enteredDate time.Time) int {
	days := enteredDate.Sub(startDate).Hours() / 24
	return int(days)
}

func readInput(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при считывании")
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
		log.Fatal(err)
	}
	return strings.TrimSpace(input)
}

//в веб
func printResult(ost float64) {
	switch ost {
	case 0:
		fmt.Println("1-я смена, в день")
	case 0.125:
		fmt.Println("2-я смена, в день")
	case 0.75, 0.25:
		fmt.Println("1й выходной")
	case 0.375, 0.875:
		fmt.Println("2й выходной")
	case 0.5:
		fmt.Println("1-я смена, в ночь")
	case 0.625:
		fmt.Println("2-я смена, в ночь")
	default:
		fmt.Println("Fail/////")
	}
}

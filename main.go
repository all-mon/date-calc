package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	emp1 := new(Employee)
	emp1.NewEmployee("Монахов", 2022, 6, 3)

	emp2 := new(Employee)
	emp2.NewEmployee("Дубинин", 2022, 6, 5)

	emp3 := new(Employee)
	emp3.NewEmployee("Тюшняков", 2022, 6, 1)

	emp4 := new(Employee)
	emp4.NewEmployee("Перехода", 2022, 6, 7)

	mapOfEmployee := make(map[string]Employee)

	mapOfEmployee[emp1.FIO] = *emp1
	mapOfEmployee[emp2.FIO] = *emp2
	mapOfEmployee[emp3.FIO] = *emp3
	mapOfEmployee[emp4.FIO] = *emp4

	enteredName := readInput("Введите фамилию:\t")

	foundEmployee, ok := mapOfEmployee[enteredName]
	if !ok {
		log.Fatal("Имени нет")
	}
	startDate := *foundEmployee.start
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
	//fmt.Printf("Раздница в днях: %v, остаток %v", differenceBetweenDates, fractional)
	fmt.Println()
	printResult(fractional)

	//2 строки что-бы консоль не закрывалась сразу
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}

//func Date(year, month, day int) *time.Time {
//	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
//	return &date
//}

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

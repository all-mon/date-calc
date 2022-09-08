package main

import (
	"bufio"
	"fmt"
	date_calc "github.com/m0n7h0ff/date-calc"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	mapOfEmployee := repository.SetEmployeeList()
	answer := getAnswer(mapOfEmployee)
	fmt.Println(answer)
	DoNotCloseTheConsole()
}

func DoNotCloseTheConsole() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
func getAnswer(mapOfEmployee map[string]date_calc.Employee) string {
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
		DoNotCloseTheConsole()
	}

	differenceBetweenDates := getDifferenceBetweenDates(startDate, parseDate)
	_, fractional := math.Modf(float64(differenceBetweenDates) / 8.0) //остаток
	fmt.Println()
	return getScheduleAnswer(fractional)
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
func getScheduleAnswer(ost float64) string {
	switch ost {
	case 0:
		return fmt.Sprint("1-я смена, в день")
	case 0.125:
		return fmt.Sprint("2-я смена, в день")
	case 0.75, 0.25:
		return fmt.Sprint("1й выходной")
	case 0.375, 0.875:
		return fmt.Sprint("2й выходной")
	case 0.5:
		return fmt.Sprint("1-я смена, в ночь")
	case 0.625:
		return fmt.Sprint("2-я смена, в ночь")
	default:
		return fmt.Sprint("Дата неверна")
	}
}

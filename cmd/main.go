package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	date_calc "github.com/m0n7h0ff/date-calc"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

type Resp struct {
	Date string
	Sch  string
}

func main() {

	//answer := getAnswer(mapOfEmployee)
	//fmt.Println(answer)
	//DoNotCloseTheConsole()yf
	r := gin.Default()
	r.GET("/api/:fio", func(context *gin.Context) {
		name := context.Param("fio")
		res := getAnswerByName(name)
		fmt.Println(res)
		context.IndentedJSON(http.StatusOK, res)
	})
	r.Run(":8080")
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

func getAnswerByName(name string) []Resp {
	mapOfEmployee := repository.GetEmployeeList()
	foundEmployee, ok := mapOfEmployee[name]
	if !ok {
		log.Fatal("Имени нет")
	}
	startDate := *foundEmployee.StartDate

	mapDate := make([]Resp, 0)
	var daysCount, year, day int
	var key time.Time
	var m time.Month
	for i := 1; i < 31; i++ {
		daysCount = 24 * i
		key = time.Now().Add(time.Duration(daysCount) * time.Hour)
		year, m, day = key.Date()
		strKey := fmt.Sprintf("%v - %v - %v", day, m, year)

		differenceBetweenDates := getDifferenceBetweenDates(startDate, key)
		_, fractional := math.Modf(float64(differenceBetweenDates) / 8.0) //остаток
		mapDate = append(mapDate, Resp{Date: strKey, Sch: getScheduleAnswer(fractional)})
	}
	return mapDate
}
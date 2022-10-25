package entities

import (
	"fmt"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
	"log"
	"math"
	"time"
)

type Employee struct {
	FIO       string `uri:"fio" binding:"required"`
	StartDate *time.Time
}

func (e *Employee) NewEmployee(name string, year, month, day int) {
	e.FIO = name
	e.StartDate = e.setDate(year, month, day)
}
func (e *Employee) setDate(year, month, day int) *time.Time {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &date
}

type Resp struct {
	Date string `json:"date" binding:"required"`
	Sch  string `json:"sch" binding:"required"`
}

func getDifferenceBetweenDates(startDate time.Time, enteredDate time.Time) int {
	days := enteredDate.Sub(startDate).Hours() / 24
	return int(days)
}

// по остатку от деления определяет смену
func getScheduleAnswer(ost float64) string {
	var shndl = map[float64]string{
		0:     "1day",
		0.125: "2day",
		0.75:  "1holiday",
		0.25:  "1holiday",
		0.375: "2holiday",
		0.875: "2holiday",
		0.5:   "1night",
		0.625: "2night",
	}
	if v, ok := shndl[ost]; ok {
		return v
	} else {
		return fmt.Sprint("Invalid date entered")
	}
	return ""
}

// Возвращает расписание на месяц по имени сторудника
func GetScheduleMonthByLname(name string) []Resp {
	mapOfEmployee := repository.GetEmployeeList()
	foundEmployee, ok := mapOfEmployee[name]
	if !ok {
		log.Println("Имени нет")
	}
	startDate := *foundEmployee.StartDate
	mapDate := make([]Resp, 0)
	var daysCount, year, day int
	var key time.Time
	var m time.Month
	for i := 1; i < 831; i++ {
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

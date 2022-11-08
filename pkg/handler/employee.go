package handler

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
)

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
		return fmt.Sprintf("Invalid input data, func getScheduleAnswer(%v)", ost)
	}
}

// Возвращает расписание на месяц по имени сторудника
func GetScheduleMonthByLname(name string) []Resp {
	mapOfEmployee := repository.EmployeePostgres.
	foundEmployee, ok := mapOfEmployee[name]
	if !ok {
		log.Println("Имени нет")
	}
	startDate := *foundEmployee.StartDate
	mapDate := make([]Resp, 0)
	var daysCount int
	var key time.Time
	for i := 1; i < 366; i++ {
		daysCount = 24 * i
		key = time.Now().Add(time.Duration(daysCount) * time.Hour)
		strKey := key.Format("2006/01/02")
		differenceBetweenDates := getDifferenceBetweenDates(startDate, key)
		_, fractional := math.Modf(float64(differenceBetweenDates) / 8.0) //остаток
		mapDate = append(mapDate, Resp{Date: strKey, Sch: getScheduleAnswer(fractional)})
	}
	return mapDate
}

func (h *Handler) getEmployeeByLastname(c *gin.Context){
	name := c.Param("fio")
	res := GetScheduleMonthByLname(name)
	c.JSON(200, res)
}

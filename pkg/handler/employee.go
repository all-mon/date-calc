package handler

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ScheduleResponse struct {
	Date         string `json:"date" binding:"required"`
	WorkingShift string `json:"working_shift" binding:"required"`
}

// страница с расписанием смен
func (h *Handler) getSchedule(c *gin.Context) {
	c.HTML(http.StatusOK, "schedule_person.tmpl.html", gin.H{
		"Title": "Hello",
	})

}

// handler, возвращает ответ ScheduleResponse struct, в формате json [Дата: х, Рабочая смена: х]
func (h *Handler) getEmployeeByLastname(c *gin.Context) {
	name := c.Param("lastname")
	res := h.GetScheduleMonthByLastName(name, c)
	c.JSON(200, res)
}

// GetScheduleMonthByLastName функция возвращает расписание на месяц по имени сторудника
func (h *Handler) GetScheduleMonthByLastName(lastname string, c *gin.Context) []ScheduleResponse {
	foundEmployee, err := h.services.Employee.GetByLastName(lastname)
	if err != nil {
		c.JSON(404, nil)
	}

	startDate := *foundEmployee.StartDate
	mapDate := make([]ScheduleResponse, 0)
	var daysCount int
	var key time.Time
	for i := 1; i < 366; i++ {
		daysCount = 24 * i
		key = time.Now().Add(time.Duration(daysCount) * time.Hour)
		strKey := key.Format("2006/01/02")
		differenceBetweenDates := getDifferenceBetweenDates(startDate, key)
		_, fractional := math.Modf(float64(differenceBetweenDates) / 8.0) //остаток
		mapDate = append(mapDate, ScheduleResponse{Date: strKey, WorkingShift: getScheduleAnswer(fractional)})
	}
	return mapDate
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

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m0n7h0ff/date-calc/cmd/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/api/:fio", func(context *gin.Context) {
		name := context.Param("fio")
		res := handlers.GetScheduleMonthByLname(name)
		context.JSON(200, res)
	})
	r.Run(":8080")
}

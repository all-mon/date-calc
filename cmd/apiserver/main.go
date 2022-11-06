package main

import (
	"log"

	"github.com/m0n7h0ff/date-calc/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(config.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	// r := gin.Default()
	// r.GET("/api/:fio", func(context *gin.Context) {
	// 	name := context.Param("fio")
	// 	res := handlers.GetScheduleMonthByLname(name)
	// 	context.JSON(200, res)
	// })
	// r.Run(":8080")
}

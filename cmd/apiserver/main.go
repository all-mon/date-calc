package main

import (
	"log"
	"github.com/m0n7h0ff/date-calc/pkg/handler"
	"github.com/m0n7h0ff/date-calc/config"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(config.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

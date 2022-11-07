package main

import (
	"log"

	"github.com/m0n7h0ff/date-calc/config"
	"github.com/m0n7h0ff/date-calc/pkg/handler"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
	"github.com/m0n7h0ff/date-calc/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services) 

	srv := new(config.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

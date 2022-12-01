package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/m0n7h0ff/date-calc/config"
	"github.com/m0n7h0ff/date-calc/pkg/handler"
	"github.com/m0n7h0ff/date-calc/pkg/repository"
	"github.com/m0n7h0ff/date-calc/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initConfig: %s", err.Error())
	}
	//точно ли нужна сторонняя библиотека для работы с переменными окружения?
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	//перенести в отдельную функцию?
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	//Внедрение зависимостей
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//настройка сервера, почитать
	srv := new(config.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
//получаем данные из файла конфигурации
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

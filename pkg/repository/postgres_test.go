package repository

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func Test_connect(t *testing.T) {
	_, err := NewPostgresDB(Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

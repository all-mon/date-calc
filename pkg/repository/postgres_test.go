package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"testing"
)

func initConfig() error {
	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func RunPostgresDB() (*sqlx.DB, error) {
	db, err := NewPostgresDB(Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "123123",
	})
	return db, err
}

func Test_connect(t *testing.T) {

	if err := initConfig(); err != nil {
		t.Errorf("error initConfig: %s", err.Error())
	}

	_, err := RunPostgresDB()
	if err != nil {
		t.Errorf("error connecting to database: %s", err.Error())
	}
}
func TestEmployeePostgres_GetByLastName(t *testing.T) {
	db, _ := RunPostgresDB()
	employeePostgres := EmployeePostgres{db: db}
	employee, err := employeePostgres.GetByLastName("Монахов")
	if err != nil {
		t.Error(err)
	}
	if employee.Lastname != "Монахов" {
		t.Error(err)
	}
}

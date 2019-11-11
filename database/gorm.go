package database

import (
	"fmt"
	"galaxy-weather/pkg/planet"
	"galaxy-weather/pkg/position"
	"galaxy-weather/pkg/weather"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
	db := Client()
	Migrate(db)
	// defer db.Close()
}

func Client() *gorm.DB {
	if db != nil {
		return db
	}
	var err error

	_ = godotenv.Load()

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	url := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", dbUsername, dbPassword, dbPort, dbName)

	db, err = gorm.Open("postgres", url)
	if err != nil {
		logrus.Error(err.Error())
		panic("failed to connect database ")
	}
	// defer db.Close()

	db.LogMode(false)
	db.DB().SetMaxIdleConns(1000)

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&planet.Planet{})
	db.AutoMigrate(&weather.Weather{})
	db.AutoMigrate(&position.Position{})
}

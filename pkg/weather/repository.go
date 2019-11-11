package weather

import (
	"galaxy-weather/database"

	"github.com/jinzhu/gorm"
)

type IWeatherRepo interface {
	GetByDay(day uint) (*Weather, error)
}

type weatherRepo struct {
	Driver *gorm.DB
}

func NewWeatherRepo() IWeatherRepo {
	return &weatherRepo{
		Driver: database.Client(),
	}
}

func (r weatherRepo) GetByDay(day uint) (*Weather, error) {
	weatherData := new(Weather)
	err := r.Driver.First(&weatherData, day).Error
	if err != nil {
		return nil, err
	}
	return weatherData, nil
}

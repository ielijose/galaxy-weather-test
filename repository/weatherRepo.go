package repository

import (
	"galaxy-weather/database"
	"galaxy-weather/model"

	"github.com/jinzhu/gorm"
)

type IWeatherRepo interface {
	GetByDay(day uint) (*model.Weather, error)
}

type weatherRepo struct {
	Driver *gorm.DB
}

func NewWeatherRepo() IWeatherRepo {
	return &weatherRepo{
		Driver: database.Client(),
	}
}

func (r weatherRepo) GetByDay(day uint) (*model.Weather, error) {
	weatherData := new(model.Weather)
	err := r.Driver.First(&weatherData, day).Error
	if err != nil {
		return nil, err
	}
	return weatherData, nil
}

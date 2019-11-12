package repository

import (
	"galaxy-weather/database"
	"galaxy-weather/model"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

type IWeatherRepo interface {
	GetByDay(day uint) (*model.Weather, error)
	Save(weather model.Weather) error
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
	err := r.Driver.Where(model.Weather{Day: day}).First(&weatherData).Error
	if err != nil {
		return nil, err
	}
	return weatherData, nil
}

func (r weatherRepo) Save(w model.Weather) error {
	err := r.Driver.
		Where(model.Weather{Day: w.Day}).
		Attrs(w).
		FirstOrCreate(&w).Error

	if err != nil {
		logrus.Errorf("[WeatherRepo.Save] (%d, %s) Error: %s", w.Day, w.WeatherType.String(), err.Error())
		return err
	}

	return nil
}

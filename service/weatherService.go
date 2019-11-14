package service

import (
	"galaxy-weather/model"
	"galaxy-weather/repository"
)

type IWeatherService interface {
	GetByDay(day uint) (*model.Weather, error)
	GetAll() (*[]model.Weather, error)
	Save(w model.Weather) error
}

type weatherService struct {
	repo repository.IWeatherRepo
}

func newWeatherService() IWeatherService {
	return &weatherService{
		repository.NewWeatherRepo(),
	}
}

var WeatherService = newWeatherService()

func (ws *weatherService) GetByDay(day uint) (*model.Weather, error) {
	data, err := ws.repo.GetByDay(day)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ws *weatherService) Save(w model.Weather) error {
	w.Weather = w.WeatherType.String()
	return ws.repo.Save(w)
}

func (ws *weatherService) GetAll() (*[]model.Weather, error) {
	return ws.repo.GetAll()
}

package service

import (
	"galaxy-weather/model"
	"galaxy-weather/repository"
)

type IWeatherService interface {
	GetByDay(day uint) (*model.Weather, error)
	PredictByDay(day uint) (*model.Weather, error)
	PredictRange(from, to uint) (*[]model.Weather, error)
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

func (ws *weatherService) PredictByDay(day uint) (*model.Weather, error) {
	return nil, nil
}

func (ws *weatherService) PredictRange(from, to uint) (*[]model.Weather, error) {
	for day := from; day < to; day++ {
		_, _ = GalaxyService.PredictWeather(day)
	}
	return nil, nil
}

func (ws *weatherService) Save(w model.Weather) error {
	w.Weather = w.WeatherType.String()
	return ws.repo.Save(w)
}

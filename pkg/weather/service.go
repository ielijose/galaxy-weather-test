package weather

import (
	"galaxy-weather/pkg/galaxy"
)

type IWeatherService interface {
	GetByDay(day uint) (*Weather, error)
	PredictByDay(day uint) (*Weather, error)
	PredictRange(from, to uint) (*[]Weather, error)
}

type weatherService struct {
	repo IWeatherRepo
}

func NewWeatherService(repo IWeatherRepo) IWeatherService {
	return &weatherService{
		repo,
	}
}

func (ws *weatherService) GetByDay(day uint) (*Weather, error) {
	data, err := ws.repo.GetByDay(day)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ws *weatherService) PredictByDay(day uint) (*Weather, error) {
	return nil, nil
}

func (ws *weatherService) PredictRange(from, to uint) (*[]Weather, error) {
	g := galaxy.NewGalaxy()

	for i := 0; i < 365; i++ {
		day := uint(i)
		w := g.PredictWeather(day)

	}
}

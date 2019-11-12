package service

import (
	"galaxy-weather/model"
	"galaxy-weather/repository"
)

type IPlanetService interface {
	Save(p model.Planet) error
}

type planetService struct {
	repo repository.IPlanetRepo
}

func newPlanetService() IPlanetService {
	return &planetService{
		repository.NewPlanetRepo(),
	}
}

var PlanetService = newPlanetService()

func (ps planetService) Save(p model.Planet) error {
	return ps.repo.Save(p)
}

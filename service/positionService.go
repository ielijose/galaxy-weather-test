package service

import (
	"galaxy-weather/model"
	"galaxy-weather/repository"
)

type IPositionService interface {
	Save(p model.Position) error
}

type positionService struct {
	repo repository.IPositionRepo
}

func newPositionService() IPositionService {
	return &positionService{
		repository.NewPositionRepo(),
	}
}

var PositionService = newPositionService()

func (ps positionService) Save(p model.Position) error {
	return ps.repo.Save(p)
}

package repository

import (
	"galaxy-weather/database"
	"galaxy-weather/model"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type IPositionRepo interface {
	Save(p model.Position) error
}

type positionRepo struct {
	Driver *gorm.DB
}

func NewPositionRepo() IPositionRepo {
	return &positionRepo{
		Driver: database.Client(),
	}
}

func (r positionRepo) Save(p model.Position) error {
	result := r.Driver.
		Where(model.Position{PlanetID: p.PlanetID, Day: p.Day}).
		Assign(model.Position{X: p.X, Y: p.Y}).
		FirstOrCreate(&p)

	if result.Error != nil {
		logrus.Errorf("[PositionRepo.Save] (%d, %d) Error: %s", p.PlanetID, p.Day, result.Error)
		return result.Error
	}
	return nil
}

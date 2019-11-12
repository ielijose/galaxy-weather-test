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
	err := r.Driver.
		Where(model.Position{PlanetID: p.PlanetID, Day: p.Day}).
		Attrs(p).
		FirstOrCreate(&p).Error

	if err != nil {
		logrus.Errorf("[PositionRepo.Save] (%d, %d) Error: %s", p.PlanetID, p.Day, err.Error())
		return err
	}
	return nil
}

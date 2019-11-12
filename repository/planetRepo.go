package repository

import (
	"galaxy-weather/database"
	"galaxy-weather/model"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

type IPlanetRepo interface {
	Save(planet model.Planet) error
}

type planetRepo struct {
	Driver *gorm.DB
}

func NewPlanetRepo() IPlanetRepo {
	return &planetRepo{
		Driver: database.Client(),
	}
}

func (pr planetRepo) Save(planet model.Planet) error {
	result := pr.Driver.Where(model.Planet{ID: planet.ID}).Attrs(planet).FirstOrCreate(&planet)

	if result.Error != nil {
		logrus.Errorf("[PlanetRepo.Save] (%s) Error: %s", planet.Name, result.Error)
		return result.Error
	}
	return nil
}

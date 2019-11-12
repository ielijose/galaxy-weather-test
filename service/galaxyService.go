package service

import "galaxy-weather/model"

var (
	ferengi = model.Planet{
		ID:              1,
		Name:            "Ferengi",
		AngularVelocity: 1,
		Distance:        500,
		Direction:       model.Clockwise,
		Radio:           80,
	}

	betasoide = model.Planet{
		ID:              2,
		Name:            "Betasoide",
		AngularVelocity: 3,
		Distance:        2000,
		Direction:       model.Clockwise,
		Radio:           30,
	}

	vulcano = model.Planet{
		ID:              3,
		Name:            "Vulcano",
		AngularVelocity: 5,
		Distance:        1000,
		Direction:       model.CounterClockwise,
		Radio:           50,
	}
)

func GetDefaultPlanets() []model.Planet {
	var planets = []model.Planet{
		ferengi,
		betasoide,
		vulcano,
	}

	return planets
}

func NewGalaxy() model.Galaxy {
	sun := model.Position{
		X: 0,
		Y: 0,
	}

	g := model.Galaxy{
		Ferengi:   ferengi,
		Betasoide: betasoide,
		Vulcano:   vulcano,
		Sun:       sun,
	}

	return g
}

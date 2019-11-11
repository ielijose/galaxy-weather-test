package galaxy

import (
	"galaxy-weather/pkg/planet"
	"galaxy-weather/pkg/position"
)

func NewGalaxy() Galaxy {
	ferengi := planet.Planet{
		Name:            "Ferengi",
		AngularVelocity: 1,
		Distance:        500,
		Direction:       planet.Clockwise,
		Radio:           80,
	}

	betasoide :=
		planet.Planet{
			Name:            "Betasoide",
			AngularVelocity: 3,
			Distance:        2000,
			Direction:       planet.Clockwise,
			Radio:           30,
		}

	vulcano := planet.Planet{
		Name:            "Vulcano",
		AngularVelocity: 5,
		Distance:        1000,
		Direction:       planet.CounterClockwise,
		Radio:           50,
	}

	sun := position.Position{
		X: 0,
		Y: 0,
	}

	g := Galaxy{
		Ferengi:   ferengi,
		Betasoide: betasoide,
		Vulcano:   vulcano,
		Sun:       sun,
	}

	return g
}

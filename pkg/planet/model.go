package planet

import (
	"galaxy-weather/pkg/position"
	"galaxy-weather/utils"
	"math"
)

const DegreesCircumference = 360

type direction uint8

const (
	Clockwise direction = iota
	CounterClockwise
)

type Planet struct {
	ID              uint      `gorm:"primary_key"`
	Name            string    `json:"name"`
	AngularVelocity uint      `json:"angular_velocity"`
	Distance        uint      `json:"distance"`
	Direction       direction `json:"direction"`
	Radio           int       `json:"radio"`
}

func (p Planet) AngularPosition(day uint) uint {
	degree := (day * p.AngularVelocity) % DegreesCircumference

	if p.Direction == Clockwise && degree > 0 {
		degree = DegreesCircumference - degree
	}

	return degree
}

func (p Planet) GetPointByDay(day uint) position.Position {
	degree := p.AngularPosition(day)

	radians := utils.DegreeToRadian(degree)

	x := math.Cos(radians) * float64(p.Distance)
	y := math.Sin(radians) * float64(p.Distance)

	return position.Position{
		X: x,
		Y: y,
	}
}

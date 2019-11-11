package position

import "math"

type Position struct {
	Day      uint    `json:"day_id" gorm:"primary_key"`
	PlanetID uint    `json:"planet_id" gorm:"primary_key"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

func (p1 *Position) DistanceOf(p2 Position) float64 {
	d := math.Sqrt(math.Pow(math.Abs(p2.X-p1.X), 2) + math.Pow(math.Abs(p2.Y-p1.Y), 2))
	return d
}

func (p1 *Position) Save() {
	
}

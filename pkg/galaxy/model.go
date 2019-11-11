package galaxy

import (
	"galaxy-weather/pkg/planet"
	"galaxy-weather/pkg/position"
	"galaxy-weather/pkg/weather"
	"galaxy-weather/utils"
	"math"
)

const MaxPerimeter = 6262.300354

type Galaxy struct {
	Ferengi   planet.Planet
	Betasoide planet.Planet
	Vulcano   planet.Planet
	Sun       position.Position
}

func (g *Galaxy) PredictWeather(day uint) (w weather.Weather) {
	w.Day = day
	w.WeatherType = weather.Unknown

	p1 := g.Ferengi.GetPointByDay(day)
	p2 := g.Betasoide.GetPointByDay(day)
	p3 := g.Vulcano.GetPointByDay(day)

	if g.AreAligned(p1, p2, p3) {
		w.WeatherType = weather.OptimalTemperature
	}

	if g.AreAlignedWithTheSun(p1, p2, p3) {
		w.WeatherType = weather.Drought
	}

	isInside, rainType := g.SunInsidePlanets(p1, p2, p3)
	if isInside {
		w.WeatherType = rainType
	}

	/* Save */
	p1.Save()

	return w
}

func (g *Galaxy) AreAligned(p3, p1, p2 position.Position) bool {
	a := p2.X - p1.X
	b := p2.Y - p1.Y
	c := p3.X - p1.X
	d := p3.Y - p1.Y
	abcd := math.Abs(a*d - b*c)

	if (abcd < a*d/15) || (abcd < b*c/15) {
		return true
	}

	return utils.FloatEquals(abcd, 0.0)
}

func (g *Galaxy) AreAlignedWithTheSun(p1, p2, p3 position.Position) bool {
	a := p2.X - p1.X
	b := p2.Y - p1.Y
	c := p3.X - p1.X
	d := p3.Y - p1.Y
	e := g.Sun.X - p1.X
	f := g.Sun.Y - p1.Y

	abcd := math.Abs(a*d - b*c)
	abef := math.Abs(a*f - b*e)

	return utils.FloatEquals(abcd, 0.0) && utils.FloatEquals(abef, 0.0)
}

func area(p1, p2, p3 position.Position) float64 {
	return math.Abs((p1.X*(p2.Y-p3.Y) + p2.X*(p3.Y-p1.Y) + p3.X*(p1.Y-p2.Y)) / 2.0)
}

func (g *Galaxy) SunInsidePlanets(p1, p2, p3 position.Position) (bool, weather.Type) {
	a := area(p1, p2, p3)

	a1 := area(g.Sun, p2, p3)
	a2 := area(p1, g.Sun, p3)
	a3 := area(p1, p2, g.Sun)

	if utils.FloatEquals(a, a1+a2+a3) && !utils.FloatEquals(a, 0.0) {
		if g.isMaxPerimeter(p1, p2, p3) {
			return true, weather.HeavyRain
		}
		return true, weather.Rain
	}

	return false, weather.Unknown
}

func (g *Galaxy) Perimeter(p1, p2, p3 position.Position) float64 {
	d1 := p1.DistanceOf(p2)
	d2 := p2.DistanceOf(p3)
	d3 := p3.DistanceOf(p1)

	return d1 + d2 + d3
}

func (g *Galaxy) isMaxPerimeter(p1, p2, p3 position.Position) bool {
	perimeter := g.Perimeter(p1, p2, p3)
	return perimeter >= MaxPerimeter
}

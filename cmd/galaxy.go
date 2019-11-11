package cmd

import (
	"galaxy-weather/pkg/galaxy"
)

func Galaxy() {
	g := galaxy.NewGalaxy()

	for i := 0; i < 10*365; i++ {
		day := uint(i)
		g.PredictWeather(day)
	}
}

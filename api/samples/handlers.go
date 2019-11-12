package samples

import (
	"fmt"
	"galaxy-weather/service"
	"net/http"

	"github.com/labstack/echo"
)

func sampleSeedHandler(c echo.Context) error {
	var response []string
	// Planets
	planets := service.GetDefaultPlanets()

	for _, planet := range planets {
		err := planet.Save()
		if err != nil {
			response = append(response, fmt.Sprintf("Planet created successfully: %s", planet.Name))
		}
	}

	g := service.NewGalaxy()

	for day := uint(0); day < 10; day++ {
		g.PredictWeather(day)
	}

	return c.JSON(http.StatusOK, response)
}

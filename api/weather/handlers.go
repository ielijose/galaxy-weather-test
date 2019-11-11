package weather

import (
	"errors"
	"galaxy-weather/database/repository"
	"galaxy-weather/pkg/weather"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func weatherByDayHandler(c echo.Context) error {
	d, _ := strconv.Atoi(c.Param("day"))
	if d < 0 {
		return errors.New("invalid day")
	}
	day := uint(d)

	weatherRepo := repository.NewWeatherRepo()
	weatherService := weather.NewWeatherService(weatherRepo)

	response, err := weatherService.GetByDay(day)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}

func predictWeatherHandler(c echo.Context) error {
	// from
	f, _ := strconv.Atoi(c.Param("from"))
	if f < 0 {
		return errors.New("invalid from day")
	}
	from := uint(f)
	// to
	t, _ := strconv.Atoi(c.Param("to"))
	if t < 0 || t < f {
		return errors.New("invalid to day")
	}
	to := uint(t)

	weatherRepo := repository.NewWeatherRepo()
	weatherService := weather.NewWeatherService(weatherRepo)

	response, err := weatherService.PredictRange(from, to)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}


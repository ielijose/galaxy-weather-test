package weather

import (
	"errors"
	"galaxy-weather/service"
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

	response, err := service.WeatherService.GetByDay(day)
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

	response, err := service.WeatherService.PredictRange(from, to)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}

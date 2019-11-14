package weather

import (
	"errors"
	"galaxy-weather/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func weatherByDayHandler(c echo.Context) error {
	d, _ := strconv.Atoi(c.QueryParam("day"))
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

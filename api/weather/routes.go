package weather

import "github.com/labstack/echo"

func Init(g *echo.Group) {
	g.GET("", weatherByDayHandler)
	g.GET("/predict", predictWeatherHandler)
}

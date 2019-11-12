package model

type Weather struct {
	Day         uint `json:"day" gorm:"primary_key"`
	WeatherType Type `json:"weather"`
}


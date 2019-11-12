package model

type Weather struct {
	Day         uint   `json:"day" gorm:"default:0;primary_key;auto_increment:false" sql:"not null"`
	WeatherType Type   `json:"-" sql:"-"`
	Weather     string `json:"weather" sql:"not null"`
}

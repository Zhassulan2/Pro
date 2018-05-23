package model

type Taxes struct {
	Id           string `gorm:"primary_key"`
	Name         string
	Category     string
	Custcategory *string
	Parentid     *string
	Rate         float64
	Ratecascade  string
	Rateorder    *int
}

package model

type Stocklevel struct {
	Id            string `gorm:"primary_key"`
	Location      string
	Product       string
	Stocksecurity float64
	Stockmaximum  float64
}

func (Stocklevel) TableName() string {
	return "stocklevel"
}

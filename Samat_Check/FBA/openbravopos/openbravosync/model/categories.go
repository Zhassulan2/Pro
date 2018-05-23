package model

type Categories struct {
	ID       string  `gorm:"primary_key" json:"id"`
	Name     string  `json:"name"`
	Parentid *string `json:"parentid"`
	Image    *string `json:"image" type:"bytea"`
}

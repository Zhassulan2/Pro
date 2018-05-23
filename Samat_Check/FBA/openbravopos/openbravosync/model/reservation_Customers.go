package model

type Reservation_Customers struct {
	Id       string `gorm:"primary_key"`
	Customer string
}

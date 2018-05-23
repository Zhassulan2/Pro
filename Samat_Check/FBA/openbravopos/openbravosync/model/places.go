package model

type Places struct {
	Id    string `gorm:"primary_key"`
	Name  string
	X     int
	Y     int
	Floor string
}

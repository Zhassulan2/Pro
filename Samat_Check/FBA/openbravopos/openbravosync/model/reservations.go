package model

import (
	"time"
)

type Reservations struct {
	Id          string `gorm:"primary_key"`
	Created     time.Time
	Datenew     time.Time
	title       string
	Chairs      int
	Isdone      string
	Description string
}

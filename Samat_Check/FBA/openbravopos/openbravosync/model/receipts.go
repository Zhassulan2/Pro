package model

import (
	"time"
)

type Receipts struct {
	Id         string `gorm:"primary_key"`
	Money      string
	Datenew    time.Time
	Attributes *string `type:"bytea"`
}

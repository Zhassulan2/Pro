package model

import (
	"time"
)

type Teambonussession struct {
	Id             string `gorm:"primary_key"`
	Tbsequence     int
	Datestart      time.Time
	Dateend        *time.Time
	Teambonussales *float32
	Teambonus      *float32
}

func (Teambonussession) TableName() string {
	return "teambonussession"
}

package model

import (
	"time"
)

type Stockdiary struct {
	Id                      string `gorm:"primary_key"`
	Datenew                 time.Time
	Reason                  int
	Location                string
	Product                 string
	Attributesetinstance_Id *string
	Units                   float64
	Price                   float64
}

func (Stockdiary) TableName() string {
	return "stockdiary"
}

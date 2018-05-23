package model

import (
	"time"
)

type Closedcash struct {
	Money        string     `gorm:"primary_key" json:"money"`
	Host         string     `json:"host"`
	Hostsequence int        `json:"hostsequence"`
	Datestart    time.Time  `json:"datestart"`
	Dateend      *time.Time `json:"dateend"`
}

func (Closedcash) TableName() string {
	return "closedcash"
}

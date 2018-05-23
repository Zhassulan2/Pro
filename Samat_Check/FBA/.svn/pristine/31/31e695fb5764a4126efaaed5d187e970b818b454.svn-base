package model

import "time"

type Customers struct {
	Id          string
	SearchKey   string    `gorm:"column:searchkey"`
	TaxId       string    `gorm:"column:taxid"`
	Name        string    `gorm:"column:name"`
	TaxCategory string    `gorm:"column:taxcategory"`
	Card        string    `gorm:"column:card"`
	MaxDebt     float32   `gorm:"column:maxdebt"`
	Address     string    `gorm:"column:address"`
	Address2    string    `gorm:"column:address2"`
	Postal      string    `gorm:"column:postal"`
	City        string    `gorm:"column:city"`
	Region      string    `gorm:"column:region"`
	Country     string    `gorm:"column:country"`
	FirstName   string    `gorm:"column:firstname"`
	LastName    string    `gorm:"column:lastname"`
	Email       string    `gorm:"column:email"`
	Phone       string    `gorm:"column:phone"`
	Phone2      string    `gorm:"column:phone2"`
	Fax         string    `gorm:"column:fax"`
	Notes       string    `gorm:"column:notes"`
	Visible     bool      `gorm:"column:visible"`
	CurDate     time.Time `gorm:"column:curdate" json:"datetime, string"`
	CurDebt     float32   `gorm:"column:curdebt"`
}

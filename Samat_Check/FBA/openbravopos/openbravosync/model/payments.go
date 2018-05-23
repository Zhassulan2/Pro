package model

type Payments struct {
	Id        string `gorm:"primary_key"`
	Receipt   string
	Payment   string
	Total     float64
	Transid   *string
	Returnmsg *string `type:"bytea"`
}

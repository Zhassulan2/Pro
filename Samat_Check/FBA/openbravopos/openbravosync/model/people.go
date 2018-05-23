package model

type People struct {
	Id          string `gorm:"primary_key"`
	Name        string
	Apppassword string
	Card        string
	Role        string
	Visible     string
	Image       *string `type:"bytea"`
}

func (People) TableName() string {
	return "people"
}

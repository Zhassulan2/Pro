package model

type Teambonussessionperson struct {
	Id                 string `gorm:"primary_key"`
	Tbs_id             string
	Person             string
	Teambonusforperson *float32
}

func (Teambonussessionperson) TableName() string {
	return "teambonussessionperson"
}

package model

type Attributeset struct {
	ID   string `gorm:"primary_key"`
	Name string
}

func (Attributeset) TableName() string {
	return "attributeset"
}

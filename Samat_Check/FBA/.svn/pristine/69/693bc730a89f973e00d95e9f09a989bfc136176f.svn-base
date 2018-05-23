package model

type Attributeuse struct {
	Id              string `gorm:"primary_key"`
	Attributeset_Id string
	Attribute_Id    string
	Lineno          *int
}

func (Attributeuse) TableName() string {
	return "attributeuse"
}

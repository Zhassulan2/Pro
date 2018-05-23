package model

type Attributeinstance struct {
	ID                      string `gorm:"primary_key"`
	Attributesetinstance_Id string
	Attribute_Id            string
	Value                   *string
}

func (Attributeinstance) TableName() string {
	return "attributeinstance"
}

package dbModel

import "github.com/go-openapi/strfmt"

//Product продукты
// swagger:model Product
type Product struct {
	ID         strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name       string       `gorm:"type:varchar;"`
	Barcode    string       `gorm:"type:varchar;"`
	Category   *Category
	CategoryID *strfmt.UUID4 `gorm:"type:uuid REFERENCES Category(Id)"`
	Staff      Staff         //Владелец
	StaffID    strfmt.UUID4  `gorm:"type:uuid REFERENCES Staff(Id)"`
}

//TableName переопределяем имя таблицы
func (Product) TableName() string {
	return "product"
}

package dbModel

import "github.com/go-openapi/strfmt"

//Supplier поставщики
// swagger:model Supplier
type Supplier struct {
	ID      strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name    string       `gorm:"type:varchar;"`
	Bin     string       `gorm:"type:varchar;"`
	Address string       `gorm:"type:varchar;"`
	Contact string       `gorm:"type:varchar;"`
	Staff   Staff
	StaffID strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"`
}

//TableName переопределяем имя таблицы
func (Supplier) TableName() string {
	return "supplier"
}

package dbModel

import "github.com/go-openapi/strfmt"

//PaymentType типы платежей
// swagger:model PaymentType
type PaymentType struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string       `gorm:"type:varchar;"`
	ShortName string       `gorm:"type:varchar;"`
}

//TableName переопределяем имя таблицы
func (PaymentType) TableName() string {
	return "paymenttype"
}

package dbModel

import "github.com/go-openapi/strfmt"

//Customer клиенты компаний
// swagger:model Customer
type Customer struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Staff     Staff
	StaffID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"`
	Address   string       `gorm:"type:varchar;"`
	Name      string       `gorm:"type:varchar;"`
	Contact   string       `gorm:"type:varchar;"`
	Iin       string       `gorm:"type:varchar;"`
	Company   Company
	CompanyID strfmt.UUID4 `gorm:"type:uuid REFERENCES Company(Id);"`
}

//TableName переопределяем имя таблицы
func (Customer) TableName() string {
	return "customer"
}

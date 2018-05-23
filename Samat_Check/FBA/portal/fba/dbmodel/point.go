package dbModel

import (
	"github.com/go-openapi/strfmt"
)

//Point точки продаж
// swagger:model Point
type Point struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Staff     Staff
	StaffID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"` //идентификатор клиента в clients_Db
	City      *City
	CityID    *strfmt.UUID4 `gorm:"type:uuid REFERENCES City(Id);default:null"`
	Address   string        `gorm:"type:varchar;"`
	Name      string        `gorm:"type:varchar;"`
	IsDeleted bool          `gorm:"type:bool;" json:"-"` //Флаг удаления записи
	ClientID  *strfmt.UUID4 `json:"-"`                   // связь с OAuth
	Company   Company
	CompanyID strfmt.UUID4 `gorm:"type:uuid REFERENCES Company(Id);"`
}

//TableName переопределяем имя таблицы
func (Point) TableName() string {
	return "point"
}

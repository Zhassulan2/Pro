package dbModel

import "github.com/go-openapi/strfmt"

//Device устройства продаж
// swagger:model Device
type Device struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point     Point
	PointID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Point(Id);"`
	Name      string       `gorm:"type:varchar;"`
	IsDeleted bool         `gorm:"type:bool;" json:"-"` //Флаг удаления записи
}

//TableName переопределяем имя таблицы
func (Device) TableName() string {
	return "device"
}

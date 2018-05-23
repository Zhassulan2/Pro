package dbModel

import "github.com/go-openapi/strfmt"

//StaffPoint привязка пользователя к точке продаж
// swagger:model StaffPoint
type StaffPoint struct {
	ID      strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point   Point
	PointID strfmt.UUID4 `gorm:"type:uuid REFERENCES Point(Id)"`
	Staff   Staff
	StaffID strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"`
}

//TableName переопределяем имя таблицы
func (StaffPoint) TableName() string {
	return "staffpoint"
}

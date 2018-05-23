package dbModel

import "github.com/go-openapi/strfmt"

// City справочник городов
// swagger:model City
type City struct {
	ID   strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string       `gorm:"type:varchar;"`
}

//TableName переопределяем имя таблицы
func (City) TableName() string {
	return "city"
}

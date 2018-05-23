package dbModel

import "github.com/go-openapi/strfmt"

//Role права пользователей(staff)
// swagger:model Role
type Role struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string       `gorm:"type:varchar;"`
	ShortName string       `gorm:"type:varchar;" json:"-"`
}

//TableName переопределяем имя таблицы
func (Role) TableName() string {
	return "role"
}

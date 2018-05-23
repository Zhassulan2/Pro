package dbModel

import "github.com/go-openapi/strfmt"

//Staff пользователь
// swagger:model Staff
type Staff struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string       `gorm:"type:varchar;"`
	Role      Role
	RoleID    strfmt.UUID4  `gorm:"type:uuid REFERENCES Role(Id)"`
	ParentID  *strfmt.UUID4 //ID владельца компании
	UserID    strfmt.UUID4  `json:"-"` //Связь с oauth
	IsDeleted bool          `gorm:"type:bool;" json:"-"`
}

//TableName переопределяем имя таблицы
func (Staff) TableName() string {
	return "staff"
}

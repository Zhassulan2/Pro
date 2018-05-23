package dbModel

import "github.com/go-openapi/strfmt"

// ActionType - список типов действий для журнала движения товаров
// swagger:model ActionType
type ActionType struct {
	// идентификатор записи
	ID strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"ID"`
	// название типа действия
	Name string `gorm:"type:varchar;unique_index:idx_actiontype_name" json:"Name"`
	// мнемоника для использования внутри кода
	ShortName string `gorm:"type:varchar;unique_index:idx_actiontype_shortname" json:"-"`
}

//TableName переопределяем имя таблицы
func (ActionType) TableName() string {
	return "actiontype"
}

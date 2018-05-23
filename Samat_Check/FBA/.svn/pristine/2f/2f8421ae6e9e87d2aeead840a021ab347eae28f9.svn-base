package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//Tax налоги
// swagger:model Tax
type Tax struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string       `gorm:"type:varchar;"`
	ValidFrom time.Time
	Value     float32
	ShortName string `gorm:"type:varchar;"`
}

//TableName переопределяем имя таблицы
func (Tax) TableName() string {
	return "tax"
}

package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//ProductAction Журнал движения товаров
// swagger:model ProductAction
type ProductAction struct {
	ID            strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point         Point
	PointID       strfmt.UUID4 `gorm:"type:uuid REFERENCES Point(Id);"`
	WorkSession   *WorkSession
	WorkSessionID *strfmt.UUID4 `gorm:"type:uuid REFERENCES WorkSession(Id);"`
	Count         int64
	Amount        float64
	ActionDate    time.Time
	Staff         Staff        // продавец (может быть и не владелец)
	StaffID       strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"`
}

//TableName переопределяем имя таблицы
func (ProductAction) TableName() string {
	return "productaction"
}

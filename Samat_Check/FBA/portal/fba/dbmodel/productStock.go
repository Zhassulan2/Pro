package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//ProductStock остатки
// swagger:model ProductStock
type ProductStock struct {
	ID          strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product     Product
	ProductID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Product(Id)"`
	Point       Point
	PointID     strfmt.UUID4 `gorm:"type:uuid REFERENCES Point(Id);"`
	Count       int64
	PriceSell   float64
	ChangeTimer time.Time `json:"-"`
}

//TableName переопределяем имя таблицы
func (ProductStock) TableName() string {
	return "productstock"
}

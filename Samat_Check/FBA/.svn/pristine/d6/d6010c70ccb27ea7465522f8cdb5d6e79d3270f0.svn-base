package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//Discount список скидок
// swagger:model Discount
type Discount struct {
	ID             strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name           string       `gorm:"type:varchar;"`
	DiscountType   DiscountType
	DiscountTypeID strfmt.UUID4 `gorm:"type:uuid REFERENCES DiscountType(Id)"`
	Startdate      time.Time    //Время начала действия скидки
	Enddate        time.Time    //Время завершения действия скидки
	Value          float64      //Значение фиксированной скидки
	Valuepercent   float64      //Значение скидки в процентах
	Staff          Staff
	StaffID        strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"` //идентификатор клиента в clients_Db
}

//TableName переопределяем имя таблицы
func (Discount) TableName() string {
	return "discount"
}

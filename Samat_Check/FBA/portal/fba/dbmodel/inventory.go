package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//Inventory журнал инвентаризаций
// swagger:model Inventory
type Inventory struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string       `gorm:"type:varchar;"`
	Startdate time.Time    //Время начала инвентаризации
	Enddate   *time.Time   //Время завершения инвентаризации
	Company   Company
	CompanyID strfmt.UUID4 `gorm:"type:uuid REFERENCES Company(Id);"`
	Staff     Staff
	StaffID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"` //идентификатор клиента в clients_Db
	Point     Point
	PointID   strfmt.UUID4
}

//TableName переопределяем имя таблицы
func (Inventory) TableName() string {
	return "inventory"
}

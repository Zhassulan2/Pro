package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//InventoryDetail детализация журнала инвентаризаций
// swagger:model InventoryDetail
type InventoryDetail struct {
	ID          strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Inventory   Inventory
	InventoryID strfmt.UUID4 `gorm:"type:uuid REFERENCES Inventory(Id)"`
	Product     Product
	ProductID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Product(Id)"`
	Staff       Staff
	StaffID     strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"` //Пользователь производивший фиксацию записи
	Count       int64        //Количество товара
	Checkdate   time.Time    //Дата записи
}

//TableName переопределяем имя таблицы
func (InventoryDetail) TableName() string {
	return "inventorydetail"
}

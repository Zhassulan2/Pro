package dbModel

import "github.com/go-openapi/strfmt"

//ProductActionDetail Детализация по операциям
// swagger:model ProductActionDetail
type ProductActionDetail struct {
	ID              strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product         Product
	ProductID       strfmt.UUID4 `gorm:"type:uuid REFERENCES Product(Id)"`
	ProductAction   ProductAction
	ProductActionID strfmt.UUID4 `gorm:"type:uuid REFERENCES ProductAction(Id)"`
	Count           int64
	Pricebuy        *float64
	Pricesell       float64
	Tax             Tax
	TaxID           strfmt.UUID4 `gorm:"type:uuid REFERENCES Tax(Id);"`
	Reference       *string      `gorm:"type:varchar;"`
	Partnumber      *string      `gorm:"type:varchar;"`
	ActionType      ActionType
	ActionTypeID    strfmt.UUID4 `gorm:"type:uuid REFERENCES ActionType(Id)"`
	PaymentType     PaymentType
	PaymentTypeID   strfmt.UUID4 `gorm:"type:uuid REFERENCES PaymentType(Id)"`
	Supplier        *Supplier
	SupplierID      *strfmt.UUID4 `gorm:"type:uuid REFERENCES Supplier(Id);"`
}

//TableName переопределяем имя таблицы
func (ProductActionDetail) TableName() string {
	return "productactiondetail"
}

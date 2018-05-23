package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//Receipt чек
// swagger:model Receipt
type Receipt struct {
	ID              strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProductAction   ProductAction
	ProductActionID strfmt.UUID4 `gorm:"type:uuid REFERENCES ProductAction(Id)"`
	Number          int64        //Номер чека
	Checkdate       time.Time    //Дата записи
	Content         string       `gorm:"type:json;"` //Содержание чека в формате JSON
	/* Структура JSON для Content
	{
		"date": "01.01.2018 00:00:00",
		"company": "Company.Name",
		"companybin": "Company.Bin",
		"companyaddress": "Company.Address",
		"pointname": "Point.Name",
		"pointaddress": "Point.Address",
		"staffname": "Staff.Name",
		"totalcount": 0.0,
		"totalamount": 0.0,
		"products": [
			{
			"name": "Product.Name",
			"count": 0.0,
			"pricesell": 0.0,
			"taxfixed": 0.0,
			"taxpercent": 0.0
			}
		]
		}
	*/
}

//TableName переопределяем имя таблицы
func (Receipt) TableName() string {
	return "receipt"
}

package dbModel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

//WorkSession смена
// swagger:model WorkSession
type WorkSession struct {
	ID          strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string       `gorm:"type:varchar;"`
	Startdate   time.Time
	Enddate     time.Time
	Isactive    bool
	Point       Point
	PointID     strfmt.UUID4 `gorm:"type:uuid REFERENCES Point(Id);"`
	Startamount float64      //Количество денег в кассе на начало сессии
	Endamount   float64      //Количество денег в кассе на закрытии сессии
	Staff       Staff
	StaffID     strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"`
}

//TableName переопределяем имя таблицы
func (WorkSession) TableName() string {
	return "worksession"
}

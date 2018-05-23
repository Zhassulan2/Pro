import { DiscountType } from "../models/discountType";
import { Staff } from "../models/staff";

export class Discount {
	ID?: string;
    Name: string;
    DiscountType?: DiscountType;
    DiscountTypeID: string;
    Startdate: string;
	Enddate: string;
	Value: number;
	Valuepercent: number;
	Staff?: Staff;
    StaffID?: string;
}

/*
type Discount struct {
	ID             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name           string    `gorm:"type:varchar;"`
	DiscountType   DiscountType
	DiscountTypeID uuid.UUID `gorm:"type:uuid REFERENCES DiscountType(Id)"`
	Startdate      time.Time //Время начала действия скидки
	Enddate        time.Time //Время завершения действия скидки
	Value          float64   //Значение фиксированной скидки
	Valuepercent   float64   //Значение скидки в процентах
	Staff          Staff
	StaffID        uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"` //идентификатор клиента в clients_Db
}
*/
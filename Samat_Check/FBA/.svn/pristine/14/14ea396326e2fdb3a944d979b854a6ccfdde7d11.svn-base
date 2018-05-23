import { Company } from "../models/company";
import { Staff } from "../models/staff";

export class Inventory {
	ID?: string;
    Name: string;
    Startdate: string;
	Enddate?: string;
	Company?: Company;
    CompanyID: string;
	Staff?: Staff;
	StaffID?: string;
	Point?: string;
	PointID?: string;
}

/*
type Inventory struct {
	ID        strfmt.UUID4 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string       `gorm:"type:varchar;"`
	Startdate time.Time    //Время начала инвентаризации
	Enddate   time.Time    //Время завершения инвентаризации
	Company   Company
	CompanyID strfmt.UUID4 `gorm:"type:uuid REFERENCES Company(Id);"`
	Staff     Staff
	StaffID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id)"` //идентификатор клиента в clients_Db
	Point     Point
	PointID   strfmt.UUID4
}
*/
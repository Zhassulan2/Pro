import { Staff } from "../models/staff";
import { City } from "../models/city";
import { Company } from "../models/company";

export class Point {
	ID?: string;
	Staff?: Staff;
	StaffID?: string;
	City?: City;
	CityID?: string;
	Address: string;
	Name: string;
	Isdeleted?: boolean;
	ClientID?: string;
	Company?: Company;
	CompanyID: string;
}

/*
type Point struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Staff     Staff
	StaffID   uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)" json:"-"` //идентификатор клиента в clients_Db
	City      City
	CityID    *uuid.UUID `gorm:"type:uuid REFERENCES City(Id);default:null" json:"-"`
	Address   string     `gorm:"type:varchar;"`
	Name      string     `gorm:"type:varchar;"`
	IsDeleted bool       `gorm:"type:bool;"` //Флаг удаления записи
	ClientID  *uuid.UUID // связь с OAuth
	Company   Company
	CompanyID uuid.UUID `gorm:"type:uuid REFERENCES Company(Id);" json:"-"`
}
*/

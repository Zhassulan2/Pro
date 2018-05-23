import { Staff } from "../models/staff";
import { Company } from "../models/company";

export class Customer {
	ID?: string;
	Staff?: Staff;
    StaffID?: string;
    Address: string;
    Name: string;
    Contact?: string;
    Iin: string;
    Company?: Staff;
    CompanyID?: string;
}
/*
type Customer struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Staff     Staff
	StaffID   uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"`
	Address   string    `gorm:"type:varchar;"`
	Name      string    `gorm:"type:varchar;"`
	Contact   string    `gorm:"type:varchar;"`
	Iin       string    `gorm:"type:varchar;"`
	Company   Company
	CompanyID uuid.UUID `gorm:"type:uuid REFERENCES Company(Id);"`
}
*/
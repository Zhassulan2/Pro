import { Staff } from "../models/staff";

export class Company {
	ID?: string;
    Name: string;
    Bin: string;
    Address: string;
	Contact?: string;
	Staff?: Staff;
    StaffID?: string;
}

/*
type Company struct {
	ID      uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name    string     `gorm:"type:varchar;"`
	Bin     string     `gorm:"type:varchar;"`
	Address string     `gorm:"type:varchar;"`
	Contact string     `gorm:"type:varchar;"`
	Staff   Staff      //`gorm:"type:uuid REFERENCES Staff(Id)"`
	StaffID *uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)" json:"-"`
}
*/
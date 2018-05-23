import { Staff } from "../models/staff";

export class Supplier {
    ID?: string;
    Name: string;
    Bin: string;
    Address: string;
    Contact: string;
    Staff?: Staff;
    StaffID?: string;
}
/*
type Supplier struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name    string    `gorm:"type:varchar;"`
	Bin     string    `gorm:"type:varchar;"`
	Address string    `gorm:"type:varchar;"`
	Contact string    `gorm:"type:varchar;"`
	Staff   Staff
	StaffID uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"`
}
*/
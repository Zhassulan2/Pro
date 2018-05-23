import { Staff } from "../models/staff";

export class Category {
	ID?: string;
	ShortName: string;
    Name: string;
	ParentID?: string;
	Staff?: Staff;
    StaffID?: string;
}

/*
type Category struct {
	ID        strfmt.UUID4  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ShortName string        `gorm:"type:varchar;unique_index:idx_category_shortname"`
	Name      string        `gorm:"type:varchar;unique_index:idx_category_name_staff"`
	ParentID  *strfmt.UUID4 `gorm:"type:uuid REFERENCES Category(Id);default:null"`
	Staff     Staff
	StaffID   strfmt.UUID4 `gorm:"type:uuid REFERENCES Staff(Id);unique_index:idx_category_name_staff"`
}
*/
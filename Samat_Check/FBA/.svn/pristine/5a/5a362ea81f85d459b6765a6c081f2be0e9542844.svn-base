import { Staff } from "../models/staff";
import { Category } from "../models/category";

export class Product {
    ID?: string;
	Name: string;
	Barcode?: string;
	Category?: Category;
	CategoryID: string;
	Staff?: Staff;
	StaffID?: string;
}

/*
type Product struct {
	ID         uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name       string    `gorm:"type:varchar;"`
	Barcode    string    `gorm:"type:varchar;"`
	Category   *Category
	CategoryID *uuid.UUID `gorm:"type:uuid REFERENCES Category(Id)"`
	Staff      Staff
	StaffID    uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"`
}
*/
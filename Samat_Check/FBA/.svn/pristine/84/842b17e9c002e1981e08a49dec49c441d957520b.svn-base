import { Product } from "../models/product";
import { Staff } from "../models/staff";
import { Inventory } from "../models/inventory";

export class InventoryDetail {
	ID?: string;
    Inventory?: Inventory;
    InventoryID: string;
	Product?: Product;
    ProductID: string;
	Staff?: Staff;
    StaffID?: string;
    Count: number;
	Checkdate?: string;
}


/*
type InventoryDetail struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Inventory   Inventory
	InventoryID uuid.UUID `gorm:"type:uuid REFERENCES Inventory(Id)"`
	Product     Product
	ProductID   uuid.UUID `gorm:"type:uuid REFERENCES Product(Id)"`
	Staff       Staff
	StaffID     uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"` //Пользователь производивший фиксацию записи
	Count       int64     //Количество товара
	Checkdate   time.Time //Дата записи
}
*/
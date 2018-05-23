import { Product } from "../models/product";
import { Point } from "../models/point";

export class ProductStock {
	ID?: string;
	Product?: Product;
	ProductID: string;
	Point?: Point;
	PointID: string;
	Count: number;
	PriceSell: number;
    ChangeTimer?: string;
}

/*
type ProductStock struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product     Product
	ProductID   uuid.UUID `gorm:"type:uuid REFERENCES Product(Id)"`
	Point       Point
	PointID     uuid.UUID `gorm:"type:uuid REFERENCES Point(Id);"`
	Count       int64
	PriceSell   float64
	ChangeTimer time.Time `json:"-"`
}
*/
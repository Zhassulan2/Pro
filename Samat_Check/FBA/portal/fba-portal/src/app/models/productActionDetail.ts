import { Product } from "../models/product";
import { ProductAction } from "../models/productAction";
import { ActionType } from "../models/actionType";
import { PaymentType } from "../models/paymentType";
import { Supplier } from "../models/supplier";
import { Tax } from "../models/tax";

export class ProductActionDetail {
	ID?: string;
	Product?: Product;
	ProductID: string;
	ProductAction?: ProductAction;
	ProductActionID: string;
	Count: number;
    Pricebuy?: number;
	Pricesell: number;
	Tax?: Tax;
    TaxID?: string;
	Reference?: string;
	Partnumber?: string;
	ActionType?: ActionType;
	ActionTypeID?: string;
	PaymentType?: PaymentType;
	PaymentTypeID?: string;
	Supplier?: Supplier;
	SupplierID?: string;
}

/*
type ProductActionDetail struct {
	ID              uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Product         Product
	ProductID       uuid.UUID `gorm:"type:uuid REFERENCES Product(Id)"`
	ProductAction   ProductAction
	ProductActionID uuid.UUID `gorm:"type:uuid REFERENCES ProductAction(Id)"`
	Count           int64
	Pricebuy        *float64
	Pricesell       float64
	Tax             Tax
	TaxID           uuid.UUID `gorm:"type:uuid REFERENCES Tax(Id);"`
	Reference       *string   `gorm:"type:varchar;"`
	Partnumber      *string   `gorm:"type:varchar;"`
	ActionType      ActionType
	ActionTypeID    uuid.UUID `gorm:"type:uuid REFERENCES ActionType(Id)"`
	PaymentType     PaymentType
	PaymentTypeID   uuid.UUID `gorm:"type:uuid REFERENCES PaymentType(Id)"`
	Supplier        *Supplier
	SupplierID      *uuid.UUID `gorm:"type:uuid REFERENCES Supplier(Id);"`
}
*/
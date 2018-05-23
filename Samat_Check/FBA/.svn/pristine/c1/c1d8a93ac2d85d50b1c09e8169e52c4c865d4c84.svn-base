import { Staff } from "../models/staff";
import { Point } from "../models/point";
import { WorkSession } from "../models/workSession";

export class ProductAction {
	ID?: string;
	Point?: Point;
	PointID: string;
	WorkSession?: WorkSession;
	WorkSessionID?: string;
	Count: number;
    Amount: number;
	ActionDate: string;
	Staff?: Staff;
	StaffID?: string;
}

/*
type ProductAction struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point         Point
	PointID       uuid.UUID `gorm:"type:uuid REFERENCES Point(Id);"`
	WorkSession   *WorkSession
	WorkSessionID *uuid.UUID `gorm:"type:uuid REFERENCES WorkSession(Id);"`
	Count         int64
	Amount        float64
	ActionDate    time.Time
	Staff         Staff
	StaffID       uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"`
}
*/
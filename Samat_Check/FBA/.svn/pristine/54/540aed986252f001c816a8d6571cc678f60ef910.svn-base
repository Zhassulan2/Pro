import { Point } from "../models/point";
import { Staff } from "../models/staff";

export class WorkSession {
	ID?: string;
	Name: string;
	Startdate?: string;
    Enddate?: string;
	Isactive:boolean;
	Point?: Point;
	PointID: string;
	Startamount: number;
	Endamount: number;
	Staff?: Staff;
	StaffID: string;
}
/*
type WorkSession struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string    `gorm:"type:varchar;"`
	Startdate   time.Time
	Enddate     time.Time
	Isactive    bool
	Point       Point
	PointID     uuid.UUID `gorm:"type:uuid REFERENCES Point(Id);"`
	Startamount float64   //Количество денег в кассе на начало сессии
	Endamount   float64   //Количество денег в кассе на закрытии сессии
	Staff       Staff
	StaffID     uuid.UUID `gorm:"type:uuid REFERENCES Staff(Id)"`
}
*/
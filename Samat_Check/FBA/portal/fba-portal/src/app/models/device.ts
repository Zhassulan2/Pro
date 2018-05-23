import { Point } from "../models/point";

export class Device {
    ID?: string;
    Point?: Point;
    PointID: string;
    Name: string;
    IsDeleted?: boolean;
}

/*
type Device struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Point     Point
	PointID   uuid.UUID `gorm:"type:uuid REFERENCES Point(Id);"`
	Name      string    `gorm:"type:varchar;"`
	IsDeleted bool      `gorm:"type:bool;"` //Флаг удаления записи
}
*/
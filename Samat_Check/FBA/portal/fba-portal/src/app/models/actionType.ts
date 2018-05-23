export class ActionType {
	ID?: string;
	Name: string;
}

/*
type ActionType struct {
	ID   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar;"`
}
*/
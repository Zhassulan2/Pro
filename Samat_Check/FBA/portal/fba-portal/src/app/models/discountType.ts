export class DiscountType {
	ID?: string;
    Name: string;
}

/*
type DiscountType struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar;"`
	ShortName string    `gorm:"type:varchar;"`
}
*/
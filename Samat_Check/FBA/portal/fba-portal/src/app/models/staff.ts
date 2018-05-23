import { Role } from "../models/role";

export class Staff {
	ID?: string;
	Name: string;
	Role?: Role;
	RoleID: string;
	ParentID: string;
	UserID?: string;
	IsDeleted?: boolean;
}

/*
type Staff struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar;"`
	Role      Role
	RoleID    uuid.UUID `gorm:"type:uuid REFERENCES Role(Id)"  json:"-"`
	ParentID  uuid.UUID //ID владельца компании ???
	UserID    uuid.UUID //Связь с oauth
	IsDeleted bool      `gorm:"type:bool;"`
}
*/
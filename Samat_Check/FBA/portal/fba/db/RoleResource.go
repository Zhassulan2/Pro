package db

import "../dbmodel"
import "fmt"
import "github.com/go-openapi/strfmt"

//RoleCreate create Role
func (dbm *DBManager) RoleCreate(role dbModel.Role) (roleRet dbModel.Role, err error) {
	if dbm.DB.NewRecord(&role) {
		err = dbm.DB.Create(&role).Error
		return role, err
	}
	return role, fmt.Errorf("%s", "запись уже существует")
}

//RoleUpdate update Role
func (dbm *DBManager) RoleUpdate(role dbModel.Role) error {
	return dbm.DB.Save(&role).Error
}

//RoleDelete delete Role
func (dbm *DBManager) RoleDelete(role dbModel.Role) error {
	return dbm.DB.Delete(&role).Error
}

//RoleGet get Role
func (dbm *DBManager) RoleGet(size, page int) (roles []dbModel.Role, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&roles)
	return
}

//RoleGetByID get by id Role
func (dbm *DBManager) RoleGetByID(id strfmt.UUID4) (role dbModel.Role, err error) {
	err = dbm.DB.Where("id = ?", id).First(&role).Error
	return
}

//RoleGetByShortName get by id Role
func (dbm *DBManager) RoleGetByShortName(shortName string) (role dbModel.Role, err error) {
	err = dbm.DB.Where("short_name = ?", shortName).First(&role).Error
	return
}

//RoleCount get count Role
func (dbm *DBManager) RoleCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Role{}).Count(&count).Error
	return
}

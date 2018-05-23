package db

import "../dbmodel"
import "../model"
import "fmt"
import "github.com/go-openapi/strfmt"

//CategoryCreate create Category
func (dbm *DBManager) CategoryCreate(category dbModel.Category, ti model.TokenInfo) (categoryRet dbModel.Category, err error) {
	uuid, err := ti.GetUserID()
	if err != nil {
		return
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		return
	}

	category.Staff = staff
	if dbm.DB.NewRecord(&category) {
		err = dbm.DB.Create(&category).Error
		return category, err
	}
	return category, fmt.Errorf("%s", "запись уже существует")
}

//CategoryUpdate update Category
func (dbm *DBManager) CategoryUpdate(category dbModel.Category) error {
	return dbm.DB.Save(&category).Error
}

//CategoryDelete delete Category
func (dbm *DBManager) CategoryDelete(category dbModel.Category) error {
	return dbm.DB.Delete(&category).Error
}

//CategoryGet get Category
func (dbm *DBManager) CategoryGet(size, page int) (categorys []dbModel.Category, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&categorys)
	for i := 0; i < len(categorys); i++ {
		categorys[i].Staff, _ = dbm.StaffGetByID(categorys[i].StaffID)
	}
	return
}

//CategoryGetByID get by id Category
func (dbm *DBManager) CategoryGetByID(id strfmt.UUID4) (category dbModel.Category, err error) {
	err = dbm.DB.Where("id = ?", id).First(&category).Error
	category.Staff, _ = dbm.StaffGetByID(category.StaffID)
	return
}

//CategoryGetByShortName get by shortName Category
func (dbm *DBManager) CategoryGetByShortName(shortName string) (category dbModel.Category, err error) {
	err = dbm.DB.Where("short_name = ?", shortName).First(&category).Error
	category.Staff, _ = dbm.StaffGetByID(category.StaffID)
	return
}

//CategoryCount get count Category
func (dbm *DBManager) CategoryCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Category{}).Count(&count).Error
	return
}

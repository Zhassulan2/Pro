package db

import "../dbmodel"
import "fmt"
import "github.com/go-openapi/strfmt"

//DiscountTypeCreate create DiscountType
func (dbm *DBManager) DiscountTypeCreate(discounttype dbModel.DiscountType) (discounttypeRet dbModel.DiscountType, err error) {
	if dbm.DB.NewRecord(&discounttype) {
		err = dbm.DB.Create(&discounttype).Error
		return discounttype, err
	}
	return discounttype, fmt.Errorf("%s", "запись уже существует")
}

//DiscountTypeUpdate update DiscountType
func (dbm *DBManager) DiscountTypeUpdate(discounttype dbModel.DiscountType) error {
	return dbm.DB.Save(&discounttype).Error
}

//DiscountTypeDelete delete DiscountType
func (dbm *DBManager) DiscountTypeDelete(discounttype dbModel.DiscountType) error {
	return dbm.DB.Delete(&discounttype).Error
}

//DiscountTypeGet get DiscountType
func (dbm *DBManager) DiscountTypeGet(size, page int) (discounttypes []dbModel.DiscountType, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&discounttypes)
	return
}

//DiscountTypeGetByID get by id DiscountType
func (dbm *DBManager) DiscountTypeGetByID(id strfmt.UUID4) (discounttype dbModel.DiscountType, err error) {
	err = dbm.DB.Where("id = ?", id).First(&discounttype).Error
	return
}

//DiscountTypeGetByShortName get by id DiscountType
func (dbm *DBManager) DiscountTypeGetByShortName(shortName string) (discounttype dbModel.DiscountType, err error) {
	err = dbm.DB.Where("short_name = ?", shortName).First(&discounttype).Error
	return
}

//DiscountTypeCount get count DiscountType
func (dbm *DBManager) DiscountTypeCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.DiscountType{}).Count(&count).Error
	return
}

package db

import "../dbmodel"
import "fmt"
import "github.com/go-openapi/strfmt"

//TaxCreate create Tax
func (dbm *DBManager) TaxCreate(tax dbModel.Tax) (taxRet dbModel.Tax, err error) {
	if dbm.DB.NewRecord(&tax) {
		err = dbm.DB.Create(&tax).Error
		return tax, err
	}
	return tax, fmt.Errorf("%s", "запись уже существует")
}

//TaxUpdate update Tax
func (dbm *DBManager) TaxUpdate(tax dbModel.Tax) error {
	return dbm.DB.Save(&tax).Error
}

//TaxDelete delete Tax
func (dbm *DBManager) TaxDelete(tax dbModel.Tax) error {
	return dbm.DB.Delete(&tax).Error
}

//TaxGet get Tax
func (dbm *DBManager) TaxGet(size, page int) (taxs []dbModel.Tax, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&taxs)
	return
}

//TaxGetById get by id Tax
func (dbm *DBManager) TaxGetById(id strfmt.UUID4) (tax dbModel.Tax, err error) {
	err = dbm.DB.Where("id = ?", id).First(&tax).Error
	return
}

//TaxCount get count Tax
func (dbm *DBManager) TaxCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Tax{}).Count(&count).Error
	return
}

//ActionTypeGetByID get by id ActionType
func (dbm *DBManager) TaxGetByShortName(shortName string) (tax dbModel.Tax, err error) {
	err = dbm.DB.Where("short_name = ?", shortName).First(&tax).Error
	return
}

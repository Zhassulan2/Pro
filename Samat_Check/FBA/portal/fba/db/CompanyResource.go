package db

import "../dbmodel"
import "../model"
import "fmt"
import "github.com/go-openapi/strfmt"

//CompanyCreate create Company
func (dbm *DBManager) CompanyCreate(company dbModel.Company, ti model.TokenInfo) (companyRet dbModel.Company, err error) {
	uuid, err := ti.GetUserID()
	if err != nil {
		return
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		return
	}

	company.Staff = staff
	if dbm.DB.NewRecord(&company) {
		err = dbm.DB.Create(&company).Error
		return company, err
	}
	return company, fmt.Errorf("%s", "запись уже существует")
}

//CompanyUpdate update Company
func (dbm *DBManager) CompanyUpdate(company dbModel.Company) error {
	return dbm.DB.Save(&company).Error
}

//CompanyDelete delete Company
func (dbm *DBManager) CompanyDelete(company dbModel.Company) error {
	return dbm.DB.Delete(&company).Error
}

//CompanyGet get Company
func (dbm *DBManager) CompanyGet(size, page int) (companys []dbModel.Company, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&companys)
	for i := 0; i < len(companys); i++ {
		companys[i].Staff, _ = dbm.StaffGetByID(companys[i].StaffID)
	}
	return
}

//CompanyGetByID get by id Company
func (dbm *DBManager) CompanyGetByID(id strfmt.UUID4) (company dbModel.Company, err error) {
	err = dbm.DB.Where("id = ?", id).First(&company).Error
	company.Staff, _ = dbm.StaffGetByID(company.StaffID)
	return
}

//CompanyCount get count Company
func (dbm *DBManager) CompanyCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Company{}).Count(&count).Error
	return
}

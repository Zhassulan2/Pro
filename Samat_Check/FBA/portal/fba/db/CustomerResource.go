package db

import "../dbmodel"
import "../model"
import "fmt"
import "github.com/go-openapi/strfmt"

//CustomerCreate create Customer
func (dbm *DBManager) CustomerCreate(customer dbModel.Customer, ti model.TokenInfo) (customerRet dbModel.Customer, err error) {
	uuid, err := ti.GetUserID()
	if err != nil {
		return 
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		return 
	}

	customer.Staff = staff

	if dbm.DB.NewRecord(&customer) {
		err = dbm.DB.Create(&customer).Error
		return customer, err
	}
	return customer, fmt.Errorf("%s", "запись уже существует")
}

//CustomerUpdate update Customer
func (dbm *DBManager) CustomerUpdate(customer dbModel.Customer) error {
	return dbm.DB.Save(&customer).Error
}

//CustomerDelete delete Customer
func (dbm *DBManager) CustomerDelete(customer dbModel.Customer) error {
	return dbm.DB.Delete(&customer).Error
}

//CustomerGet get Customer
func (dbm *DBManager) CustomerGet(size, page int) (customers []dbModel.Customer, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&customers)
	for i := 0; i < len(customers); i++ {
		customers[i].Company, _ = dbm.CompanyGetByID(customers[i].CompanyID)
		customers[i].Staff, _ = dbm.StaffGetByID(customers[i].StaffID)
	}
	return
}

//CustomerGetByID get by id Customer
func (dbm *DBManager) CustomerGetByID(id strfmt.UUID4) (customer dbModel.Customer, err error) {
	err = dbm.DB.Where("id = ?", id).First(&customer).Error
	customer.Company, _ = dbm.CompanyGetByID(customer.CompanyID)
	customer.Staff, _ = dbm.StaffGetByID(customer.StaffID)
	return
}

//CustomerCount get count Customer
func (dbm *DBManager) CustomerCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Customer{}).Count(&count).Error
	return
}

package db

import "../dbmodel"
import "../model"
import "fmt"
import "github.com/go-openapi/strfmt"

//SupplierCreate create Supplier
func (dbm *DBManager) SupplierCreate(supplier dbModel.Supplier, ti model.TokenInfo) (supplierRet dbModel.Supplier, err error) {
	uuid, err := ti.GetUserID()
	if err != nil {
		return 
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		return 
	}

	supplier.Staff = staff

	if dbm.DB.NewRecord(&supplier) {
		err = dbm.DB.Create(&supplier).Error
		return supplier, err
	}
	return supplier, fmt.Errorf("%s", "запись уже существует")
}

//SupplierUpdate update Supplier
func (dbm *DBManager) SupplierUpdate(supplier dbModel.Supplier) error {
	return dbm.DB.Save(&supplier).Error
}

//SupplierDelete delete Supplier
func (dbm *DBManager) SupplierDelete(supplier dbModel.Supplier) error {
	return dbm.DB.Delete(&supplier).Error
}

//SupplierGet get Supplier
func (dbm *DBManager) SupplierGet(size, page int) (suppliers []dbModel.Supplier, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&suppliers)
	for i := 0; i < len(suppliers); i++ {
		suppliers[0].Staff, _ = dbm.StaffGetByID(suppliers[0].StaffID)
	}
	return
}

//SupplierGetByID get by id Supplier
func (dbm *DBManager) SupplierGetByID(id strfmt.UUID4) (supplier dbModel.Supplier, err error) {
	err = dbm.DB.Where("id = ?", id).First(&supplier).Error
	supplier.Staff, _ = dbm.StaffGetByID(supplier.StaffID)
	return
}

//SupplierCount get count Supplier
func (dbm *DBManager) SupplierCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Supplier{}).Count(&count).Error
	return
}

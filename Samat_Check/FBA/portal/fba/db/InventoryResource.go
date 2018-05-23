package db

import "../dbmodel"
import "../model"
import "fmt"
import "github.com/go-openapi/strfmt"

//InventoryCreate create Inventory
func (dbm *DBManager) InventoryCreate(inventory dbModel.Inventory, ti model.TokenInfo) (inventoryRet dbModel.Inventory, err error) {
	uuid, err := ti.GetUserID()
	if err != nil {
		return
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		return
	}

	inventory.Staff = staff
	if dbm.DB.NewRecord(&inventory) {
		err = dbm.DB.Create(&inventory).Error
		return inventory, err
	}
	return inventory, fmt.Errorf("%s", "запись уже существует")
}

//InventoryUpdate update Inventory
func (dbm *DBManager) InventoryUpdate(inventory dbModel.Inventory) error {
	return dbm.DB.Save(&inventory).Error
}

//InventoryDelete delete Inventory
func (dbm *DBManager) InventoryDelete(inventory dbModel.Inventory) error {
	return dbm.DB.Delete(&inventory).Error
}

//InventoryGet get Inventory
func (dbm *DBManager) InventoryGet(size, page int) (inventorys []dbModel.Inventory, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&inventorys)
	for i := 0; i < len(inventorys); i++ {
		inventorys[i].Company, _ = dbm.CompanyGetByID(inventorys[i].CompanyID)
		inventorys[i].Staff, _ = dbm.StaffGetByID(inventorys[i].StaffID)
		inventorys[i].Point, _ = dbm.PointGetByID(inventorys[i].PointID)
	}
	return
}

//InventoryGetByID get by id Inventory
func (dbm *DBManager) InventoryGetByID(id strfmt.UUID4) (inventory dbModel.Inventory, err error) {
	err = dbm.DB.Where("id = ?", id).First(&inventory).Error
	inventory.Company, _ = dbm.CompanyGetByID(inventory.CompanyID)
	inventory.Staff, _ = dbm.StaffGetByID(inventory.StaffID)
	inventory.Point, _ = dbm.PointGetByID(inventory.PointID)
	return
}

//InventoryCount get count Inventory
func (dbm *DBManager) InventoryCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Inventory{}).Count(&count).Error
	return
}

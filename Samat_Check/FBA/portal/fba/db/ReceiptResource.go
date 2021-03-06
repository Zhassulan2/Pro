package db

import "../dbmodel"
import "fmt"
import "github.com/satori/go.uuid"

//ReceiptCreate create Receipt
func (dbm *DBManager) ReceiptCreate(receipt dbModel.Receipt) (receiptRet dbModel.Receipt, err error) {
	if dbm.DB.NewRecord(&receipt) {
		err = dbm.DB.Create(&receipt).Error
		return receipt, err
	}
	return receipt, fmt.Errorf("%s", "запись уже существует")
}

//ReceiptUpdate update Receipt
func (dbm *DBManager) ReceiptUpdate(receipt dbModel.Receipt) error {
	return dbm.DB.Save(&receipt).Error
}

//ReceiptDelete delete Role
func (dbm *DBManager) ReceiptDelete(receipt dbModel.Receipt) error {
	return dbm.DB.Delete(&receipt).Error
}

//ReceiptGet get Role
func (dbm *DBManager) ReceiptGet(size, page int) (receipts []dbModel.Receipt, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&receipts)
	return
}

//ReceiptGetByID get by id Role
func (dbm *DBManager) ReceiptGetByID(id uuid.UUID) (receipt dbModel.Receipt, err error) {
	err = dbm.DB.Where("id = ?", id).First(&receipt).Error
	return
}

//ReceiptCount get count Role
func (dbm *DBManager) ReceiptCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Receipt{}).Count(&count).Error
	return
}

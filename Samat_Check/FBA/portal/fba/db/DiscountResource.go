package db

import "../dbmodel"
import "../model"
import "fmt"
import "github.com/go-openapi/strfmt"

//DiscountCreate create Discount
func (dbm *DBManager) DiscountCreate(discount dbModel.Discount, ti model.TokenInfo) (discountRet dbModel.Discount, err error) {
	uuid, err := ti.GetUserID()
	if err != nil {
		return 
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		return 
	}

	discount.Staff = staff
	if dbm.DB.NewRecord(&discount) {
		err = dbm.DB.Create(&discount).Error
		return discount, err
	}
	return discount, fmt.Errorf("%s", "запись уже существует")
}

//DiscountUpdate update Discount
func (dbm *DBManager) DiscountUpdate(discount dbModel.Discount) error {
	return dbm.DB.Save(&discount).Error
}

//DiscountDelete delete Discount
func (dbm *DBManager) DiscountDelete(discount dbModel.Discount) error {
	return dbm.DB.Delete(&discount).Error
}

//DiscountGet get Discount
func (dbm *DBManager) DiscountGet(size, page int) (discounts []dbModel.Discount, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&discounts)
	for i := 0; i < len(discounts); i++ {
		discounts[i].DiscountType, _ = dbm.DiscountTypeGetByID(discounts[i].DiscountTypeID)
		discounts[i].Staff, _ = dbm.StaffGetByID(discounts[i].StaffID)
	}
	return
}

//DiscountGetByID get by id Discount
func (dbm *DBManager) DiscountGetByID(id strfmt.UUID4) (discount dbModel.Discount, err error) {
	err = dbm.DB.Where("id = ?", id).First(&discount).Error
	discount.DiscountType, _ = dbm.DiscountTypeGetByID(discount.DiscountTypeID)
	discount.Staff, _ = dbm.StaffGetByID(discount.StaffID)
	return
}

//DiscountCount get count Discount
func (dbm *DBManager) DiscountCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Discount{}).Count(&count).Error
	return
}

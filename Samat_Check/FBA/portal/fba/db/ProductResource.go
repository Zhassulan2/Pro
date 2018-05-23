package db

import (
	"fmt"
	"strconv"

	"../dbmodel"
	"../model"

	"github.com/go-openapi/strfmt"
)

//ProductCreate create Product
func (dbm *DBManager) ProductCreate(product dbModel.Product, ti model.TokenInfo) (productRet dbModel.Product, err error) {

	uuid, err := ti.GetUserID()
	if err != nil {
		return
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		return
	}

	product.Staff = staff

	if dbm.DB.NewRecord(&product) {
		err = dbm.DB.Create(&product).Error
		return product, err
	}
	return product, fmt.Errorf("%s", "запись уже существует")
}

//ProductUpdate update Product
func (dbm *DBManager) ProductUpdate(product dbModel.Product) error {
	return dbm.DB.Save(&product).Error
}

//ProductDelete delete Product
func (dbm *DBManager) ProductDelete(product dbModel.Product) error {
	return dbm.DB.Delete(&product).Error
}

//ProductGet get Product
func (dbm *DBManager) ProductGet(params map[string][]string) (products []dbModel.Product, err error) {
	page := 1
	size := 10
	if len(params["page"]) > 0 {
		page, _ = strconv.Atoi(params["page"][0])
	}

	if len(params["size"]) > 0 {
		size, _ = strconv.Atoi(params["size"][0])
	}

	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&products)
	for i := 0; i < len(products); i++ {
		if products[i].CategoryID != nil {
			cat, _ := dbm.CategoryGetByID(*products[i].CategoryID)
			products[i].Category = &cat
		}

		products[i].Staff, _ = dbm.StaffGetByID(products[i].StaffID)
	}
	return
}

//ProductGetByID get by id Product
func (dbm *DBManager) ProductGetByID(id strfmt.UUID4) (product dbModel.Product, err error) {
	err = dbm.DB.Where("id = ?", id).First(&product).Error

	if product.CategoryID != nil {
		cat, _ := dbm.CategoryGetByID(*product.CategoryID)
		product.Category = &cat
	}

	product.Staff, _ = dbm.StaffGetByID(product.StaffID)
	return
}

//ProductCount get count Product
func (dbm *DBManager) ProductCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Product{}).Count(&count).Error
	return
}

//ProductSearch поиск или выдача спика продуктов
func (dbm *DBManager) ProductSearch(params map[string][]string, ti model.TokenInfo) (products []dbModel.Product, err error) {
	uuid, err := ti.GetUserID()
	if err != nil {
		return nil, err
	}

	staff, err := dbm.StaffGetByUserID(uuid)
	if err != nil {
		fmt.Println("STAFF", uuid)
		return nil, err
	}
	var user dbModel.Staff

	if staff.ParentID == nil /*.String() != "00000000-0000-0000-0000-000000000000"*/ {
		user, err = dbm.StaffGetByID(*staff.ParentID)
	} else {
		user = staff
	}

	page := 1
	size := 10
	var name, barcode string
	if len(params["page"]) > 0 {
		page, _ = strconv.Atoi(params["page"][0])
	}

	if len(params["size"]) > 0 {
		size, _ = strconv.Atoi(params["size"][0])
	}

	if len(params["name"]) > 0 {
		name = params["name"][0]
	}

	if len(params["barcode"]) > 0 {
		barcode = params["barcode"][0]
	}

	if name == "" && barcode == "" {
		dbm.DB.Where("staff_id = ?", user.ID).Limit(size).Order("id asc").Offset((page - 1) * size).Find(&products)
		//return
	}

	if name != "" {
		dbm.DB.Where("staff_id = ?", user.ID).Where("Name like ?", "%"+name+"%").Limit(size).Order("id asc").Offset((page - 1) * size).Find(&products)
		//return
	}

	if barcode != "" {
		dbm.DB.Where("staff_id = ?", user.ID).Where("Barcode like ?", "%"+barcode+"%").Limit(size).Order("id asc").Offset((page - 1) * size).Find(&products)
		//return
	}

	for i := 0; i < len(products); i++ {
		if products[i].CategoryID != nil {
			cat, _ := dbm.CategoryGetByID(*products[i].CategoryID)
			products[i].Category = &cat
		}

		products[i].Staff, _ = dbm.StaffGetByID(products[i].StaffID)
	}
	return
}

func (dbm *DBManager) ProductAveragePrice(productID, pointID strfmt.UUID4) (averagePrice float64) {
	actionType, err := dbm.ActionTypeGetByShortName("acceptance")

	query := `select AVG(pricesell) from productactiondetail
	where product_action_id in (select id from productaction where point_id = ?)
	  and action_type_id = ?
	  and product_id = ?`
	err = dbm.DB.Raw(query, pointID, actionType.ID, productID).Scan(&averagePrice).Error
	if err != nil {
		return 0
	}
	return
}

func (dbm *DBManager) ProductLastPrice(productID, pointID strfmt.UUID4) (lastPrice float64) {
	var productStock dbModel.ProductStock
	err := dbm.DB.Where("product_id = ? and point_id = ?", productID, pointID).First(&productStock).Error
	if err != nil {
		return 0
	}
	return productStock.PriceSell
}

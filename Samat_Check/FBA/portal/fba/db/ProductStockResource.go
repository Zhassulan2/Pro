package db

import (
	"errors"
	"../dbmodel"
	"../model"
	"fmt"

	"github.com/go-openapi/strfmt"
)

//ProductStockCreate create ProductStock
func (dbm *DBManager) ProductStockCreate(productstock dbModel.ProductStock) (productstockRet dbModel.ProductStock, err error) {
	if dbm.DB.NewRecord(&productstock) {
		err = dbm.DB.Create(&productstock).Error
		return productstock, err
	}
	return productstock, fmt.Errorf("%s", "запись уже существует")
}

//ProductStockUpdate update ProductStock
func (dbm *DBManager) ProductStockUpdate(productstock dbModel.ProductStock) error {
	return dbm.DB.Save(&productstock).Error
}

//ProductStockDelete delete ProductStock
func (dbm *DBManager) ProductStockDelete(productstock dbModel.ProductStock) error {
	return dbm.DB.Delete(&productstock).Error
}

//ProductStockGet get ProductStock
func (dbm *DBManager) ProductStockGet(size, page int, p model.Parameters) (productstocks []dbModel.ProductStock, err error) {

	q, single := p.ToQuery()
	fmt.Println(single, q)
	if single {
		err = errors.New("method not single")
		return
	}
	if q == "" {
		dbm.DB.Limit(size).Offset((page - 1) * size).Find(&productstocks)
	} else {
		dbm.DB.Where(q).Limit(size).Offset((page - 1) * size).Find(&productstocks)
	}
	//return
	//dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&productstocks)
	for i := 0; i < len(productstocks); i++ {
		productstocks[i].Point, _ = dbm.PointGetByID(productstocks[i].PointID)
		productstocks[i].Product, _ = dbm.ProductGetByID(productstocks[i].ProductID)
	}
	return
}

//ProductStockSingleGet get ProductStock
func (dbm *DBManager) ProductStockSingleGet(size, page int, p model.Parameters) (productstock dbModel.ProductStock, err error) {
	q, single := p.ToQuery()
	if !single {
		err = errors.New("method is single")
		return
	}
	if q == "" {
		err = errors.New("no parameters")
		return
	}
	err = dbm.DB.Where(q).Limit(size).Offset((page - 1) * size).First(&productstock).Error
	productstock.Point, _ = dbm.PointGetByID(productstock.PointID)
	productstock.Product, _ = dbm.ProductGetByID(productstock.ProductID)
	return
}

//ProductStockGetByPointID get by id ProductStock
func (dbm *DBManager) ProductStockGetByPointID(pointID strfmt.UUID4, size, page int) (productstocks []dbModel.ProductStock, err error) {
	//err = dbm.DB.Where("point_id = ?", pointID).First(&productstock).Error
	dbm.DB.Where("point_id = ?", pointID).Limit(size).Offset((page - 1) * size).Find(&productstocks)
	for i := 0; i < len(productstocks); i++ {
		productstocks[i].Point, _ = dbm.PointGetByID(productstocks[i].PointID)
		productstocks[i].Product, _ = dbm.ProductGetByID(productstocks[i].ProductID)
	}
	return
}

//ProductStockSearchGet get by id ProductStock
func (dbm *DBManager) ProductStockSearchGet(pointID strfmt.UUID4, size, page int, productname, productbarcode string) (productstocks []dbModel.ProductStock, err error) {
	var ids []strfmt.UUID4

	if productname != "" {
		productname = "%" + productname + "%"
	}
	if productbarcode != "" {
		productbarcode = "%" + productbarcode + "%"
	}

	//if productname != "" || productbarcode != "" {
	dbm.DB.Where("name LIKE ?", productname).Or("barcode LIKE ?", productbarcode).Find(&dbModel.Product{}).Pluck("id", &ids)
	//}
	if len(ids) > 0 {
		dbm.DB.Where("point_id = ?", pointID).Where("product_id in (?)", ids).Limit(size).Offset((page - 1) * size).Find(&productstocks)
	} else {
		dbm.DB.Where("point_id = ?", pointID).Where("product_id in (?)", ids).Limit(size).Offset((page - 1) * size).Find(&productstocks)
	}
	for i := 0; i < len(productstocks); i++ {
		productstocks[i].Point, _ = dbm.PointGetByID(productstocks[i].PointID)
		productstocks[i].Product, _ = dbm.ProductGetByID(productstocks[i].ProductID)
	}
	return
}

//ProductStockGetByID get by id ProductStock
func (dbm *DBManager) ProductStockGetByID(id strfmt.UUID4) (productstock dbModel.ProductStock, err error) {
	err = dbm.DB.Where("id = ?", id).First(&productstock).Error
	productstock.Point, _ = dbm.PointGetByID(productstock.PointID)
	productstock.Product, _ = dbm.ProductGetByID(productstock.ProductID)
	return
}

//ProductStockCount get count ProductStock
func (dbm *DBManager) ProductStockCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.ProductStock{}).Count(&count).Error
	return
}

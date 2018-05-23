package db

import (
	"encoding/json"
	"../dbmodel"
	"../model"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
)

//ImportProduct create ImportProduct
func (dbm *DBManager) ImportProduct(pointID strfmt.UUID4, importProducts []model.ImportProduct, ti model.TokenInfo) (pa dbModel.ProductAction, err error) {
	//Завести продукты
	var products []model.ImportProduct

	for _, p := range importProducts {
		product := p
		importProducts = importProducts[1:]

		for _, pp := range importProducts {
			if p.BarCode == pp.BarCode && strings.ToUpper(p.Name) == strings.ToUpper(pp.Name) {
				product.Count += pp.Count
				remove(importProducts, pp)
			}

			if p.BarCode == pp.BarCode && strings.ToUpper(p.Name) != strings.ToUpper(pp.Name) {
				//error
				b, _ := json.Marshal(p)
				errorMessage := "[" + string(b)
				b, _ = json.Marshal(pp)
				errorMessage += string(b) + "]"
				err = model.NewError(errorMessage)
				return
			}

			if p.BarCode != pp.BarCode && strings.ToUpper(p.Name) == strings.ToUpper(pp.Name) {
				//error
				b, _ := json.Marshal(p)
				errorMessage := "[" + string(b)
				b, _ = json.Marshal(pp)
				errorMessage += string(b) + "]"
				err = model.NewError(errorMessage)
				return
			}
		}

		products = append(products, product)
	}

	// =================================== CREATE PRODUCT ACTION ===================================
	pa.ActionDate = time.Now()
	pa.PointID = pointID
	pa, err = dbm.ProductActionCreate(pa, ti)
	if err != nil {
		return
	}
	// =================================== END PRODUCT ACTION ===================================
	var productActionDetails []dbModel.ProductActionDetail
	actionType, err := dbm.ActionTypeGetByShortName("acceptance")
	paymentType, err := dbm.PaymentTypeGetByShortName("nopay")
	tax, err := dbm.TaxGetByShortName("notax")
	for _, p := range products {
		category, err := dbm.CategoryGetByShortName(p.CategoryName)
		if err != nil {
			return pa, err
		}
		dbProduct := dbModel.Product{Name: p.Name, Barcode: p.BarCode, CategoryID: &category.ID}
		dbProduct, err = dbm.ProductCreate(dbProduct, ti)

		productActionDetails = append(productActionDetails, dbModel.ProductActionDetail{ProductID: dbProduct.ID,
			ProductActionID: pa.ID,
			Count:           p.Count,
			Pricebuy:        &p.PriceBuy,
			Pricesell:       p.PriceSell,
			TaxID:           tax.ID,
			ActionTypeID:    actionType.ID,
			PaymentTypeID:   paymentType.ID})
	}

	productActionDetails, err = dbm.ProductActionDetailsCreate(productActionDetails)
	if err != nil {
		dbm.ProductActionDelete(pa)
		return
	}

	pa, _ = dbm.ProductActionGetByID(pa.ID)

	return
}

func remove(list []model.ImportProduct, item model.ImportProduct) []model.ImportProduct {
	for i, v := range list {
		if v == item {
			copy(list[i:], list[i+1:])
			list = list[:len(list)-1]
			return list
		}
	}
	return list
}

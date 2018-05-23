package DataBaseManager

import (
	"net/http"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TaxcategoriesResource структура для связывания
type TaxcategoriesResource struct {
}

//Create add model.Taxcategories
func (tr *TaxcategoriesResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Taxcategories{}))
}

//Delete remove model.Taxcategories
func (tr *TaxcategoriesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Taxcategories{}))
}

func (tr *TaxcategoriesResource) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var taxcategories []model.Taxcategories
	db.Find(&taxcategories)
	c.JSON(http.StatusOK, taxcategories)
}

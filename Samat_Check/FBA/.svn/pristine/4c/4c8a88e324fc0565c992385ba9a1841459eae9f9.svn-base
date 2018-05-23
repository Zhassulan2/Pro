package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//TaxcustcategoriesResource структура для связывания
type TaxcustcategoriesResource struct{}

//Create add model.Taxcustcategories
func (tr *TaxcustcategoriesResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Taxcustcategories{}))
}

//Delete remove model.Taxcustcategories
func (tr *TaxcustcategoriesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Taxcustcategories{}))
}

package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//TaxesResource структура для связывания
type TaxesResource struct{}

//Create add model.Taxes
func (tr *TaxesResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Taxes{}))
}

//Delete remove model.Taxes
func (tr *TaxesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Taxes{}))
}

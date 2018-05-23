package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//TaxlinesResource структура для связывания
type TaxlinesResource struct {
}

//Create add model.Taxlines
func (tr *TaxlinesResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Taxlines{}))
}

//Delete remove model.Taxlines
func (tr *TaxlinesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Taxlines{}))
}

package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//CustomerResource структура для связывания
type CustomerResource struct{}

//Create add model.Customers
func (tr *CustomerResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Customers{}))
}

//Delete remove model.Customers
func (tr *CustomerResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Customers{}))
}

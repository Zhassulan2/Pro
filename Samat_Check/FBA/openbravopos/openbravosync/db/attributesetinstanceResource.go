package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//AttributesetinstanceResource структура для связывания
type AttributesetinstanceResource struct{}

//Create add model.Attributesetinstance
func (tr *AttributesetinstanceResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Attributesetinstance{}))
}

//Delete remove model.Attributesetinstance
func (tr *AttributesetinstanceResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Attributesetinstance{}))
}

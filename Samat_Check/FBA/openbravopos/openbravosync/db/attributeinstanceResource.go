package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//AttributeinstanceResource структура для связывания
type AttributeinstanceResource struct{}

//Create add model.Attributeinstance
func (tr *AttributeinstanceResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Attributeinstance{}))
}

//Delete remove model.Attributeinstance
func (tr *AttributeinstanceResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Attributeinstance{}))
}

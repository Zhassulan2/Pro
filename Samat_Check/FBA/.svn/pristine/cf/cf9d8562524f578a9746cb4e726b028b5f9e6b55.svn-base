package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//AttributevalueResource структура для связывания
type AttributevalueResource struct{}

//Create add model.Attributevalue
func (tr *AttributevalueResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Attributevalue{}))
}

//Delete remove model.Attributevalue
func (tr *AttributevalueResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Attributevalue{}))
}

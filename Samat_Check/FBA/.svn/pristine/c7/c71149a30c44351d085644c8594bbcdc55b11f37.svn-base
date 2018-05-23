package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//AttributeResource структура для связывания
type AttributeResource struct{}

//Create add model.Attribute
func (tr *AttributeResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Attribute{}))
}

//Delete remove model.Attribute
func (tr *AttributeResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Attribute{}))
}

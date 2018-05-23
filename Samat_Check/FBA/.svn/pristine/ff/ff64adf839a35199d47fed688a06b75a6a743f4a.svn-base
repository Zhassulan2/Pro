package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//AttributesetResource структура для связывания
type AttributesetResource struct{}

//Create add model.Attributeset
func (tr *AttributesetResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Attributeset{}))
}

//Delete remove model.Attributeset
func (tr *AttributesetResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Attributeset{}))
}

package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//AttributeuseResource структура для связывания
type AttributeuseResource struct{}

//Create add model.Attributeuse
func (tr *AttributeuseResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Attributeuse{}))
}

//Delete remove model.Attributeuse
func (tr *AttributeuseResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Attributeuse{}))
}

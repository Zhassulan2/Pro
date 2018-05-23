package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//TeambonussessionResource структура для связывания
type TeambonussessionResource struct{}

//Create add model.Teambonussession
func (tr *TeambonussessionResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Teambonussession{}))
}

//Delete remove model.Teambonussession
func (tr *TeambonussessionResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Teambonussession{}))
}

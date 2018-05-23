package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//TeambonussessionpersonResource структура для связывания
type TeambonussessionpersonResource struct{}

//Create add model.Teambonussessionperson
func (tr *TeambonussessionpersonResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Teambonussessionperson{}))
}

//Delete remove model.Teambonussessionperson
func (tr *TeambonussessionpersonResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Teambonussessionperson{}))
}

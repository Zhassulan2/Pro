package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//TeambonusconfighostsResource структура для связывания
type TeambonusconfighostsResource struct {
}

//Create add model.Teambonusconfighosts
func (tr *TeambonusconfighostsResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Teambonusconfighosts{}))
}

//Delete remove model.Teambonusconfighosts
func (tr *TeambonusconfighostsResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Teambonusconfighosts{}))
}

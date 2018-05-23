package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//ThirdpartiesResource структура для связывания
type ThirdpartiesResource struct {
}

//Create add model.Thirdparties
func (tr *ThirdpartiesResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Thirdparties{}))
}

//Delete remove model.Thirdparties
func (tr *ThirdpartiesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Thirdparties{}))
}

package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//StocklevelResource структура для связывания
type StocklevelResource struct {
}

//Create add model.Stocklevel
func (tr *StocklevelResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Stocklevel{}))
}

//Delete remove model.Stocklevel
func (tr *StocklevelResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Stocklevel{}))
}

package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//LocationsResource структура для связывания
type LocationsResource struct{}

//Create add model.Locations
func (tr *LocationsResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Locations{}))
}

//Delete remove model.Locations
func (tr *LocationsResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Locations{}))
}

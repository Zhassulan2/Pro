package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//PlacesResource структура для связывания
type PlacesResource struct {
}

//Create add model.Places
func (tr *PlacesResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Places{}))
}

//Delete remove model.Places
func (tr *PlacesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Places{}))
}

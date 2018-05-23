package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//ReservationsResource структура для связывания
type ReservationsResource struct{}

//Create add model.Reservations
func (tr *ReservationsResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Reservations{}))
}

//Delete remove model.Reservations
func (tr *ReservationsResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Reservations{}))
}

package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//ReservationCustomersResource структура для связывания
type ReservationCustomersResource struct{}

//Create add model.Reservation_Customers
func (tr *ReservationCustomersResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Reservation_Customers{}))
}

//Delete remove model.Reservation_Customers
func (tr *ReservationCustomersResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Reservation_Customers{}))
}

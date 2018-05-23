package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//TicketResource структура для связывания
type TicketResource struct {
}

//Create добавление
func (tr *TicketResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Tickets{}))
}

//Delete удаление
func (tr *TicketResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Tickets{}))
}

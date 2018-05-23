package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//TicketLineResource структура для связывания
type TicketLineResource struct {
}

//Create add model.Ticketlines
func (tr *TicketLineResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.TicketlinesInput
	var valueDB model.Ticketlines

	if c.Bind(&value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}
	value.Attributes = ToString(value.Attributes)

	valueDB.Ticket = value.Id.Ticket
	valueDB.Line = value.Id.Line
	valueDB.Product = value.Product
	valueDB.Attributesetinstance_id = value.Attributesetinstance_id
	valueDB.Units = value.Units
	valueDB.Price = value.Price
	valueDB.Taxid = value.Taxid
	valueDB.Attributes = value.Attributes

	if err := db.Save(&valueDB).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, valueDB)
}

//Delete remove model.Ticketlines
func (tr *TicketLineResource) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	params := strings.Split(id, ":")

	var value model.Ticketlines
	db.Where("ticket = ? AND line = ?", params[0], params[1]).First(&value)
	db.Delete(&value)
	c.JSON(201, "")
}

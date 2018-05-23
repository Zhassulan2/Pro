package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//PaymentsResource структура для связывания
type PaymentsResource struct {
}

//Create add model.Payments
func (tr *PaymentsResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Payments

	if c.Bind(&value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}

	value.Returnmsg = ToString(value.Returnmsg)

	if err := db.Save(&value).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, value)
}

//Delete remove model.Payments
func (tr *PaymentsResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Payments{}))
}

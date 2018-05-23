package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//ReceiptResource структура для связывания
type ReceiptResource struct{}

//Create add model.Receipts
func (tr *ReceiptResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Receipts

	if c.Bind(&value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}

	value.Attributes = ToString(value.Attributes)

	if err := db.Save(&value).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, value)
}

//Delete remove model.Receipts
func (tr *ReceiptResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Receipts{}))
}

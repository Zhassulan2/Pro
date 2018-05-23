package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//SharedticketsResource структура для связывания
type SharedticketsResource struct {
}

//Create add model.Sharedtickets
func (tr *SharedticketsResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Sharedtickets

	if c.Bind(&value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}
	value.Content = ToString(value.Content)

	if err := db.Save(&value).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, value)
}

//Delete remove model.Sharedtickets
func (tr *SharedticketsResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Sharedtickets{}))
}

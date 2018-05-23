package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//FloorsResource структура для связывания
type FloorsResource struct {
}

//Create add model.Floors
func (tr *FloorsResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Floors

	if c.Bind(&value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}

	value.Image = ToString(value.Image)

	if err := db.Save(&value).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, value)
}

//Delete remove model.Floors
func (tr *FloorsResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Floors{}))
}

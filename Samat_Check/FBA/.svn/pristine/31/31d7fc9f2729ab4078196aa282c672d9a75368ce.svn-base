package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//ResourcesResource структура для связывания
type ResourcesResource struct{}

//Create add model.Resources
func (tr *ResourcesResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Resources

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

//Delete remove model.Resources
func (tr *ResourcesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Resources{}))
}

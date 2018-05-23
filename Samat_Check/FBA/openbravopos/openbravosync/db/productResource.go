package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//ProductResource структура для связывания
type ProductResource struct {
}

//Create add model.Products
func (tr *ProductResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Products

	if c.Bind(&value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}
	value.Image = ToString(value.Image)
	value.Attributes = ToString(value.Attributes)

	if err := db.Save(&value).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, value)
}

//Delete remove model.Products
func (tr *ProductResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Products{}))
}

func (tr *ProductResource) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value []model.Products
	db.Find(&value)
	c.JSON(201, value)
}

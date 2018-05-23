package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//CategoriesResource структура для связывания
type CategoriesResource struct{}

//Create add model.Categories
func (tr *CategoriesResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Categories

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

//Delete remove model.Categories
func (tr *CategoriesResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Categories{}))
}

func (tr *CategoriesResource) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value []model.Categories

	db.Find(&value)
	c.JSON(201, value)
}

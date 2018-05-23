package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Products_CatResource структура для связывания
type Products_CatResource struct{}

//Create add model.Products_Cat
func (tr *Products_CatResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Products_Cat{}))
}

//Delete remove model.Products_Cat
func (tr *Products_CatResource) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var value model.Products_Cat
	db.Where("product = ?", id).First(&value)
	db.Delete(&value)
	c.JSON(201, "")
}

func (tr *Products_CatResource) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value []model.Products_Cat
	db.Find(&value)
	c.JSON(201, value)
}

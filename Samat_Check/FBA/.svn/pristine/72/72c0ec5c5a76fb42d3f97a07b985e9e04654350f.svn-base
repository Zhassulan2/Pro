package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Products_ComResource структура для связывания
type Products_ComResource struct{}

//Create add model.Products_Com
func (tr *Products_ComResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Products_Com{}))
}

//Delete remove model.Products_Com
func (tr *Products_ComResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Products_Com{}))
}

func (tr *Products_ComResource) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value []model.Products_Com
	db.Find(&value)
	c.JSON(201, value)
}

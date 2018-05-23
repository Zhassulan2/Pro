package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//StockcurrentResource структура для связывания
type StockcurrentResource struct{}

//Create add model.Stockcurrent
func (tr *StockcurrentResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Stockcurrent{}))
}

//Create remove model.Stockcurrent
func (tr *StockcurrentResource) Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var value model.Stockcurrent
	db.Where("stockdiary_id = ?", id).First(&value)
	db.Delete(&value)
	c.JSON(201, "")
}

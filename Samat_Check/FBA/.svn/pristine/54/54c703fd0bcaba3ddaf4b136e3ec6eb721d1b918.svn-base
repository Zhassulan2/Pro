package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//ClosedcashResource структура для связывания
type ClosedcashResource struct {
}

//Create add model.Closedcash
func (tr *ClosedcashResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Closedcash{}))
}

//Delete remove model.Closedcash
func (tr *ClosedcashResource) Delete(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var value model.Closedcash
	db.Where("money = ?", id).First(&value)
	db.Delete(&value)
	c.JSON(201, "")
}

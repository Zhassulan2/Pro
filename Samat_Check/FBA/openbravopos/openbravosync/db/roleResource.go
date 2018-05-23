package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//RoleResource структура для связывания
type RoleResource struct{}

//Create add model.Roles
func (tr *RoleResource) Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var value model.Roles

	if c.Bind(&value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}
	value.Permissions = ToString(value.Permissions)

	if err := db.Save(&value).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, value)
}

//Delete remove model.Roles
func (tr *RoleResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Roles{}))
}

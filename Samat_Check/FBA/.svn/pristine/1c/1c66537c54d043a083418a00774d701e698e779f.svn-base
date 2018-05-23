package DataBaseManager

import (
	"encoding/hex"
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// структура для связывания
type DataBaseManager struct {
}

//GetDb получения строки
func GetDb(settings model.DataBaseConnect) (gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", settings.Host, settings.Port, settings.User, settings.DBName, settings.Password)
	db, err := gorm.Open("postgres", connectionString)
	db.LogMode(false)
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(80)
	db.DB().SetMaxIdleConns(9)
	return *db, err
}

//ToString decoding java to db
func ToString(input *string) *string {
	if input != nil {
		buf := string(*input)
		b, _ := hex.DecodeString(buf[2:])
		x := string(b)
		return &x
	}
	return input
}

//DefaultCreate стандартное создание объекта
func DefaultCreate(c *gin.Context, typeObject reflect.Type) {
	var db *gorm.DB
	db = c.MustGet("db").(*gorm.DB)
	value := reflect.New(typeObject).Interface()

	if c.Bind(value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		c.Request.Body.Close()
		c.Request.Close = true
		return
	}
	c.Request.Body.Close()
	if err := db.Save(value).Error; err != nil {
		fmt.Println("ERROR CREATE", err)
		db.Rollback()
		c.Request.Body.Close()
		c.Request.Close = true
		c.AbortWithStatus(404)
		//c.JSON(404, err)
		return
	}
	db.Commit()
	c.JSON(201, value)
}

//DefaultDelete стандартное удаление объекта
func DefaultDelete(c *gin.Context, typeObject reflect.Type) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	value := reflect.New(typeObject).Interface()
	db.Where("id = ?", id).First(&value)
	db.Delete(&value)
	c.JSON(201, "")
}

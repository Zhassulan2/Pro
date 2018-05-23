package DataBaseManager

import (
	"net/http"
	"openbravosync/model"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

//StockdiaryResource структура для связывания
type StockdiaryResource struct{}

//Create добавление в журнал
func (tr *StockdiaryResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Stockdiary{}))
}

//Delete удаление из журнала
func (tr *StockdiaryResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Stockdiary{}))
}

//GetP Получение операций с указанной даты
func (tr *StockdiaryResource) GetP(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	date := c.Param("date")
	var stockdiarys []model.Stockdiary
	db.Where("Datenew = ?", date).Find(&stockdiarys)
	c.JSON(http.StatusOK, stockdiarys)
}

//Get Получение операций за весь период
func (tr *StockdiaryResource) Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	date := c.Param("date")
	t, _ := time.Parse("01-08-2016T22:19:32:1932", date)
	var stockdiarys []model.Stockdiary
	if date == "" {
		db.Find(&stockdiarys)
	} else {
		db.Where("Datenew > ?", t).Find(&stockdiarys)
	}
	c.JSON(http.StatusOK, stockdiarys)
}

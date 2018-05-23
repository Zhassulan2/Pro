package DataBaseManager

import (
	"openbravosync/model"
	"reflect"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

//ApplicationResource структура для связывания
type ApplicationResource struct {
	db *gorm.DB
}

//Create add model.Applications
func (*ApplicationResource) Create(c *gin.Context) {
	DefaultCreate(c, reflect.TypeOf(model.Applications{}))
}

//Delete remove model.Applications
func (*ApplicationResource) Delete(c *gin.Context) {
	DefaultDelete(c, reflect.TypeOf(model.Applications{}))
}

//GetSync получить данные для отправки на сервер
func (ar *ApplicationResource) GetSync() (applications []model.Applications, err error) {
	err = ar.db.Where("sync = 0").Find(&applications).Error
	return
}

//SetSync установить признак синхранизации
func (ar *ApplicationResource) SetSync(application model.Applications) error {
	err := ar.db.Exec("update application set sync = 1 where id = '" + application.Id + "';").Error
	return err
}

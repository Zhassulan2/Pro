package db

import "../dbmodel"
import "fmt"

import "github.com/go-openapi/strfmt"

//CityCreate create City
func (dbm *DBManager) CityCreate(city dbModel.City) (cityRet dbModel.City, err error) {
	if dbm.DB.NewRecord(&city) {
		err = dbm.DB.Create(&city).Error
		return city, err
	}
	return city, fmt.Errorf("%s", "запись уже существует")
}

//CityUpdate update City
func (dbm *DBManager) CityUpdate(city dbModel.City) error {
	return dbm.DB.Save(&city).Error
}

//CityDelete delete City
func (dbm *DBManager) CityDelete(city dbModel.City) error {
	return dbm.DB.Delete(&city).Error
}

//CityGet get City
func (dbm *DBManager) CityGet(size, page int) (citys []dbModel.City, count int, err error) {
	//var count int
	dbm.DB.Limit(size).Order("name asc").Offset((page - 1) * size).Find(&citys).Count(&count)
	return
}

//CityGetByID get by id City
func (dbm *DBManager) CityGetByID(id strfmt.UUID4) (city dbModel.City, err error) {
	err = dbm.DB.Where("id = ?", id).First(&city).Error
	return
}

//CityCount get count City
func (dbm *DBManager) CityCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.City{}).Count(&count).Error
	return
}
package db

import "../dbmodel"
import "fmt"
import "github.com/go-openapi/strfmt"

//StaffPointCreate create StaffPoint
func (dbm *DBManager) StaffPointCreate(staffpoint dbModel.StaffPoint) (staffpointRet dbModel.StaffPoint, err error) {
	if dbm.DB.NewRecord(&staffpoint) {
		err = dbm.DB.Create(&staffpoint).Error
		return staffpoint, err
	}
	return staffpoint, fmt.Errorf("%s", "запись уже существует")
}

//StaffPointUpdate update StaffPoint
func (dbm *DBManager) StaffPointUpdate(staffpoint dbModel.StaffPoint) error {
	return dbm.DB.Save(&staffpoint).Error
}

//StaffPointDelete delete StaffPoint
func (dbm *DBManager) StaffPointDelete(staffpoint dbModel.StaffPoint) error {
	return dbm.DB.Delete(&staffpoint).Error
}

//StaffPointGet get StaffPoint
func (dbm *DBManager) StaffPointGet(size, page int) (staffpoints []dbModel.StaffPoint, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&staffpoints)
	for i := 0; i < len(staffpoints); i++ {
		staffpoints[i].Point, _ = dbm.PointGetByID(staffpoints[i].PointID)
		staffpoints[i].Staff, _ = dbm.StaffGetByID(staffpoints[i].StaffID)
	}
	return
}

//StaffPointGetByID get by id StaffPoint
func (dbm *DBManager) StaffPointGetByID(id strfmt.UUID4) (staffpoint dbModel.StaffPoint, err error) {
	err = dbm.DB.Where("id = ?", id).First(&staffpoint).Error
	staffpoint.Point, _ = dbm.PointGetByID(staffpoint.PointID)
	staffpoint.Staff, _ = dbm.StaffGetByID(staffpoint.StaffID)
	return
}

//StaffPointCount get count StaffPoint
func (dbm *DBManager) StaffPointCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.StaffPoint{}).Count(&count).Error
	return
}

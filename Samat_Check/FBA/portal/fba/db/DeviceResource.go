package db

import "../dbmodel"
import "fmt"
import "github.com/go-openapi/strfmt"

//DeviceCreate create Device
func (dbm *DBManager) DeviceCreate(device dbModel.Device) (deviceRet dbModel.Device, err error) {
	if dbm.DB.NewRecord(&device) {
		err = dbm.DB.Create(&device).Error
		return device, err
	}
	return device, fmt.Errorf("%s", "запись уже существует")
}

//DeviceUpdate update Device
func (dbm *DBManager) DeviceUpdate(device dbModel.Device) error {
	return dbm.DB.Save(&device).Error
}

//DeviceDelete delete Device
func (dbm *DBManager) DeviceDelete(device dbModel.Device) error {
	return dbm.DB.Delete(&device).Error
}

//DeviceGet get Device
func (dbm *DBManager) DeviceGet(size, page int) (devices []dbModel.Device, err error) {
	dbm.DB.Limit(size).Order("id asc").Offset((page - 1) * size).Find(&devices)
	for i := 0; i < len(devices); i++ {
		devices[i].Point, _ = dbm.PointGetByID(devices[i].PointID)
	}
	return
}

//DeviceGetByID get by id Device
func (dbm *DBManager) DeviceGetByID(id strfmt.UUID4) (device dbModel.Device, err error) {
	err = dbm.DB.Where("id = ?", id).First(&device).Error
	device.Point, _ = dbm.PointGetByID(device.PointID)
	return
}

//DeviceCount get count Device
func (dbm *DBManager) DeviceCount() (count int, err error) {
	err = dbm.DB.Model(&dbModel.Device{}).Count(&count).Error
	return
}

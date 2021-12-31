package model

import (
	"FrontAccess/src/service"
)

type Device struct {
	ID         string `gorm:""column:id;primary_key"`
	TenantId   string `gorm:""column:tenant_id;index:idx_info_device_01"`
	ApplyId    string `gorm:""column:apply_id;index:idx_info_device_02"`
	DeviceName string `gorm:""column:device_name"`
	Imei       string `gorm:""column:imei;index:idx_info_device_03"`
	Imsi       string `gorm:""column:imsi"`
	Sim        string `gorm:""column:sim"`
	DeviceId   string `gorm:""column:device_id;index:idx_info_device_05"`
}

//设置表名，默认是结构体的名的复数形式
func (Device) TableName() string {
	return "info_device"
}

func GetDeviceId(Imei string) (user Device, err error) {
	var DeviceId Device

	ret := service.MysqlClient.Raw("select device_id from info_device where imei = ?", Imei).Scan(&DeviceId)
	return DeviceId, ret.Error
}

func UpdateDeviceId(Imei string, DeviceName string) (err error) {

	ret := service.MysqlClient.Model(&Device{}).Where("imei = ?", Imei).Update("device_name", DeviceName)
	return ret.Error
}

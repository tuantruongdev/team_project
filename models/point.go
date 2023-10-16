package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

const DEFAULT_POINTS = 0

type Device struct {
	Id         int32     `json:"id" gorm:"primarykey"`
	DeviceId   string    `json:"device_id" gorm:"unique"`
	Point      int32     `json:"point" gorm:"point"`
	LastUpdate time.Time `json:"last_update" gorm:"default:current_timestamp"`
}

func (*Device) TableName() string {
	return "fake_msg_points"
}

func InsertDevice(db *gorm.DB, point *Device) bool {
	if err := db.Create(&point).Error; err != nil {
		fmt.Println(point, err.Error())
		return false
	}
	return true
}

func GetDevice(db *gorm.DB, deviceId string) (Device, error) {
	var device Device
	if err := db.Where("device_id = ? ", deviceId).First(&device).Error; err != nil {
		return device, err
	}
	return device, nil
}

func UpdatePoint(db *gorm.DB, device *Device) bool {
	if err := db.Model(&device).Where("device_id = ? ", device.DeviceId).Update("point", device.Point).Error; err != nil {
		return false
	}
	return true
}

func GetLeaderboard(db *gorm.DB) ([]Device, error) {
	var results []Device
	err := db.Model(&results).Select("id, point").Order("point DESC").Limit(100).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

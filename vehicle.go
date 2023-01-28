package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Id         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId     *uuid.UUID     `gorm:"type:uuid;column:user_id" json:"userId"`
	User       User           `gorm:"foreignKey:UserId"`
	DeviceId   *uuid.UUID     `gorm:"type:uuid;column:device_id;unique" json:"deviceId"`
	Device     Device         `gorm:"foreignKey:DeviceId"`
	Type       string         `gorm:"column:type;" json:"type"`
	VMake      string         `gorm:"column:make;" json:"make"`
	VModel     string         `gorm:"column:model;" json:"model"`
	RegNumber  string         `gorm:"column:reg_number;" json:"regNumber"`
	Attributes datatypes.JSON `gorm:"type:jsonb;column:attributes;" json:"attributes"`
	Nickname   string         `gorm:"column:nickname;" json:"nickname"`
	Odometer   string         `gorm:"column:odometer;" json:"odometer"`
	Mileage    string         `gorm:"column:mileage;" json:"mileage"`
	SpeedLimit string         `gorm:"column:speed_limit;" json:"speedLimit"`
	FuelLevel  string         `gorm:"column:fuel_level;" json:"fuelLevel"`
}

func (Vehicle) TableName() string {
	return "vehicles"
}

func (u Vehicle) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"type":       u.Type,
		"make":       u.VMake,
		"model":      u.VModel,
		"regNumber":  u.RegNumber,
		"nickname":   u.Nickname,
		"odometer":   u.Odometer,
		"mileage":    u.Mileage,
		"speedLimit": u.SpeedLimit,
		"createdAt":  u.CreatedAt,
	}
	return payload
}

func (u Vehicle) Json(preload bool) map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"type":       u.Type,
		"make":       u.VMake,
		"model":      u.VModel,
		"regNumber":  u.RegNumber,
		"attributes": u.Attributes,
		"nickname":   u.Nickname,
		"odometer":   u.Odometer,
		"mileage":    u.Mileage,
		"speedLimit": u.SpeedLimit,
		"createdAt":  u.CreatedAt,
	}

	if preload {
		payload["user"] = u.User.ShortJson()
		payload["device"] = u.Device.Json()
	}

	return payload
}

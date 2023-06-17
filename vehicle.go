package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Id            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId        *uuid.UUID     `gorm:"type:uuid;column:user_id" json:"userId"`
	User          User           `gorm:"foreignKey:UserId"`
	DeviceId      *uuid.UUID     `gorm:"type:uuid;column:device_id;unique" json:"deviceId"`
	Device        Device         `gorm:"foreignKey:DeviceId"`
	Type          int16          `gorm:"column:type;default:1" json:"type"`
	VMake         string         `gorm:"column:make;" json:"make"`
	VModel        string         `gorm:"column:model;" json:"model"`
	EngineNumber  string         `gorm:"column:engine_number;" json:"engineNumber"`
	ChassisNumber string         `gorm:"column:chassis_number;" json:"chassisNumber"`
	RegNumber     string         `gorm:"column:reg_number;" json:"regNumber"`
	Attributes    datatypes.JSON `gorm:"type:jsonb;column:attributes;" json:"attributes"`
	Nickname      string         `gorm:"column:nickname;" json:"nickname"`
	Odometer      sql.NullInt16  `gorm:"column:odometer;" json:"odometer"`
	Mileage       sql.NullInt16  `gorm:"column:mileage;" json:"mileage"`
	SpeedLimit    sql.NullInt16  `gorm:"column:speed_limit;" json:"speedLimit"`
	FuelLevel     sql.NullInt16  `gorm:"column:fuel_level;" json:"fuelLevel"`
	Status        int16          `gorm:"column:status;" json:"status"`
	Active        bool           `gorm:"column:active;default:true" json:"active"`
}

func (Vehicle) TableName() string {
	return "vehicles"
}

func (u Vehicle) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"type":      u.Type,
		"make":      u.VMake,
		"model":     u.VModel,
		"regNumber": u.RegNumber,
		"nickname":  u.Nickname,
		"status":    u.Status,
		"active":    u.Active,
		"createdAt": u.CreatedAt,
	}
	return payload
}

func (u Vehicle) Json(preload bool) map[string]interface{} {
	payload := map[string]interface{}{
		"id":            u.Id,
		"type":          u.Type,
		"make":          u.VMake,
		"model":         u.VModel,
		"regNumber":     u.RegNumber,
		"engineNumber":  u.EngineNumber,
		"chassisNumber": u.ChassisNumber,
		"attributes":    u.Attributes,
		"nickname":      u.Nickname,
		"odometer":      u.Odometer.Int16,
		"mileage":       u.Mileage.Int16,
		"speedLimit":    u.SpeedLimit.Int16,
		"fuelLevel":     u.FuelLevel.Int16,
		"status":        u.Status,
		"active":        u.Active,
		"createdAt":     u.CreatedAt,
	}

	if preload {
		payload["user"] = u.User.ShortJson()
		payload["device"] = u.Device.ShortJson()
	}

	return payload
}

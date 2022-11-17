package models

import (
	"time"

	"github.com/google/uuid"
)

type DeviceLocationHistory struct {
	Id                     uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	DeviceId               *uuid.UUID `gorm:"type:uuid;column:device_id;unique" json:"deviceId"`
	Device                 Device     `gorm:"foreignKey:DeviceId"`
	Position               string     `gorm:"column:position;" json:"position"`
	Speed                  float32    `gorm:"column:speed;" json:"speed"`
	Altitude               float32    `gorm:"column:altitude;" json:"altitude"`
	Provider               int16      `gorm:"column:provider;" json:"provider"`
	ReadAt                 time.Time  `gorm:"column:read_at;" json:"readAt"`
	IgnitionStatus         bool       `gorm:"column:ignition_status;" json:"ignitionStatus"`
	Heading                float32    `gorm:"column:heading;" json:"heading"`
	InputVoltage           float32    `gorm:"column:input_voltage;" json:"inputVoltage"`
	InternalBatteryVoltage float32    `gorm:"column:internal_battery_voltage;" json:"internalBatteryVoltage"`
	GsmSignalStrength      int16      `gorm:"column:gsm_signal_strength;" json:"gsmSignalStrength"`
	CreatedAt              time.Time  `gorm:"column:created_at;" json:"createdAt"`
}

func (DeviceLocationHistory) TableName() string {
	return "device_location_histories"
}

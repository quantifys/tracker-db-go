package models

import (
	"time"

	"github.com/google/uuid"
)

type DeviceLocationHistory struct {
	Id                     uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	DeviceId               *uuid.UUID `gorm:"type:uuid;column:device_id;unique" json:"deviceId"`
	Device                 Device     `gorm:"foreignKey:DeviceId"`
	Position               GeoJson    `gorm:"column:position;" json:"position"`
	Speed                  float64    `gorm:"column:speed;" json:"speed"`
	Altitude               float64    `gorm:"column:altitude;" json:"altitude"`
	Provider               int16      `gorm:"column:provider;" json:"provider"`
	ReadAt                 time.Time  `gorm:"column:read_at;" json:"readAt"`
	IgnitionStatus         bool       `gorm:"column:ignition_status;" json:"ignitionStatus"`
	Heading                float64    `gorm:"column:heading;" json:"heading"`
	InputVoltage           float64    `gorm:"column:input_voltage;" json:"inputVoltage"`
	InternalBatteryVoltage float64    `gorm:"column:internal_battery_voltage;" json:"internalBatteryVoltage"`
	GsmSignalStrength      int16      `gorm:"column:gsm_signal_strength;" json:"gsmSignalStrength"`
	NetworkOperator        string     `gorm:"column:network_operator;" json:"networkOperator"`
	PacketString           string     `gorm:"column:packet_string;" json:"packetString"`
	CreatedAt              time.Time  `gorm:"column:created_at;" json:"createdAt"`
}

func (DeviceLocationHistory) TableName() string {
	return "device_location_histories"
}

func (u DeviceLocationHistory) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"altitude":     u.Altitude,
		"position":    u.Position,
		"speed":     u.Speed,
		"provider": u.Provider,
		"readAt":   u.ReadAt,
		"ignitionStatus":   u.IgnitionStatus,
		"heading":   u.Heading,
		"networkOperator":   u.NetworkOperator,
	}
	return payload
}
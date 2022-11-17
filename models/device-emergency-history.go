package models

import (
	"time"

	"github.com/google/uuid"
)

type DeviceEmergencyHistory struct {
	DeviceId  uuid.UUID `gorm:"column:device_id" json:"deviceId"`
	Position  GeoJson   `gorm:"column:position" json:"position"`
	Distance  float32   `gorm:"column:distance" json:"distance"`
	Speed     float32   `gorm:"column:speed" json:"speed"`
	Altitude  float32   `gorm:"column:altitude" json:"altitude"`
	ReadAt    time.Time `gorm:"column:read_at" json:"readAt"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	Device    Device    `gorm:"foreignKey:DeviceId"`
}

func (DeviceEmergencyHistory) TableName() string {
	return "device_emergency_histories"
}

func (d DeviceEmergencyHistory) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"deviceId":  d.DeviceId,
		"position":  d.Position,
		"distance":  d.Distance,
		"speed":     d.Speed,
		"altitude":  d.Altitude,
		"readAt":    d.ReadAt.Format(time.RFC3339),
		"createdAt": d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

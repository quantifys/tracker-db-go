package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeofenceMember struct {
	gorm.Model
	Id         uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	VehicleId  uuid.UUID       `gorm:"column:vehicle_id" json:"vehicleId"`
	DeviceId   uuid.UUID       `gorm:"column:device_id" json:"deviceId"`
	GeofenceId uuid.UUID       `gorm:"column:geofence_id" json:"geofenceId"`
	Vehicle    Vehicle         `gorm:"foreignKey:VehicleId"`
	Device     Device          `gorm:"foreignKey:DeviceId"`
	Geofence   VehicleGeofence `gorm:"foreignKey:GeofenceId"`
}

func (GeofenceMember) TableName() string {
	return "geofence_members"
}

func (d GeofenceMember) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         d.Id,
		"vehicleId":  d.VehicleId,
		"deviceId":   d.DeviceId,
		"geofenceId": d.GeofenceId,
		"createdAt":  d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

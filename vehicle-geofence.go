package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleGeofence struct {
	gorm.Model
	Id       uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string         `gorm:"column:name" json:"name"`
	Geometry GeoJsonPolygon `gorm:"column:geometry" json:"geometry"`
	UserId   uuid.UUID      `gorm:"column:user_id" json:"userId"`
	User     User           `gorm:"foreignKey:UserId"`
}

func (VehicleGeofence) TableName() string {
	return "vehicle_geofences"
}

func (d VehicleGeofence) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        d.Id,
		"name":      d.Name,
		"userId":    d.UserId,
		"geometry":  d.Geometry,
		"createdAt": d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

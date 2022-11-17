package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GeofenceViolation struct {
	gorm.Model
	Id           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	VehicleId    uuid.UUID `gorm:"column:vehicle_id" json:"vehicleId"`
	StartedAt    time.Time `gorm:"column:started_at" json:"startedAt"`
	EndedAt      time.Time `gorm:"column:ended_at" json:"endedAt"`
	IsViolated   bool      `gorm:"column:is_violated" json:"isViolated"`
	LastLocation GeoJson   `gorm:"column:serial_no" json:"serialNo"`
}

func (GeofenceViolation) TableName() string {
	return "geofence_violations"
}

func (d GeofenceViolation) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":           d.Id,
		"vehicleId":    d.VehicleId,
		"startedAt":    d.StartedAt.Format(time.RFC3339),
		"endedAt":      d.EndedAt.Format(time.RFC3339),
		"isViolated":   d.IsViolated,
		"lastLocation": d.LastLocation,
		"createdAt":    d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

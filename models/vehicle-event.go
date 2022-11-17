package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleEvent struct {
	gorm.Model
	Id           uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	VehicleId    uuid.UUID    `gorm:"column:vehicle_id" json:"vehicleId"`
	StartedAt    time.Time    `gorm:"column:started_at" json:"startedAt"`
	EndedAt      sql.NullTime `gorm:"column:ended_at" json:"endedAt"`
	LastLocation GeoJson      `gorm:"column:last_location" json:"lastLocation"`
	IsComplete   bool         `gorm:"column:is_complete;default:false" json:"isComplete"`
	EventType    int32        `gorm:"type:integer;column:event_type" json:"eventType"`
	Vehicle      Vehicle      `gorm:"foreignKey:VehicleId"`
}

func (VehicleEvent) TableName() string {
	return "vehicle_events"
}

func (d VehicleEvent) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":           d.Id,
		"vehicleId":    d.VehicleId,
		"startedAt":    d.StartedAt,
		"endedAt":      d.EndedAt,
		"lastLocation": d.LastLocation,
		"isComplete":   d.IsComplete,
		"eventType":    d.EventType,
		"createdAt":    d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

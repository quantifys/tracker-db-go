package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleDriver struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId    uuid.UUID `gorm:"column:user_id" json:"userId"`
	DriverId  uuid.UUID `gorm:"column:driver_id" json:"driverId"`
	VehicleId uuid.UUID `gorm:"column:vehicle_id" json:"vehicleId"`
	User      User      `gorm:"foreignKey:UserId"`
	Driver    User      `gorm:"foreignKey:DriverId"`
	Vehicle   Vehicle   `gorm:"foreignKey:VehicleId"`
}

func (VehicleDriver) TableName() string {
	return "vehicle_drivers"
}

func (d VehicleDriver) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        d.Id,
		"userId":    d.UserId,
		"driverId":  d.DriverId,
		"vehicleId": d.VehicleId,
		"createdAt": d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

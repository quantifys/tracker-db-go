package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FuelEntry struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Amount    float32   `gorm:"column:amount" json:"amount"`
	Quantity  float32   `gorm:"column:quantity" json:"quantity"`
	FuelType  int32     `gorm:"column:fuel_type" json:"fuelType"`
	FilledOn  time.Time `gorm:"column:filled_on" json:"filledOn"`
	VehicleId uuid.UUID `gorm:"column:vehicle_id" json:"vehicleId"`
	Vehicle   Vehicle   `gorm:"foreignKey:VehicleId"`
}

func (FuelEntry) TableName() string {
	return "fuel_entries"
}

func (d FuelEntry) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        d.Id,
		"amount":    d.Amount,
		"quantity":  d.Quantity,
		"fuelType":  d.FuelType,
		"filledOn":  d.FilledOn,
		"vehicleId": d.VehicleId,
		"createdAt": d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleDocument struct {
	gorm.Model
	Id              uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	DocumentType    int32        `gorm:"column:document_type" json:"documentType"`
	ExpiringOn      sql.NullTime `gorm:"column:expiring_on" json:"expiringOn"`
	RenewalReminder time.Time    `gorm:"column:renewal_reminder" json:"renewalReminder"`
	Url             string       `gorm:"column:url" json:"url"`
	VehicleId       uuid.UUID    `gorm:"column:vehicle_id" json:"vehicleId"`
	IsPublic        bool         `gorm:"column:is_public;default:true" json:"isPublic"`
	Vehicle         Vehicle      `gorm:"foreignKey:VehicleId"`
}

func (VehicleDocument) TableName() string {
	return "vehicle_documents"
}

func (d VehicleDocument) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":              d.Id,
		"documentType":    d.DocumentType,
		"expiringOn":      nil,
		"renewalReminder": d.RenewalReminder.Format(time.RFC3339),
		"url":             d.Url,
		"vehicleId":       d.VehicleId,
		"createdAt":       d.CreatedAt.Format(time.RFC3339),
	}
	if d.ExpiringOn.Valid {
		payload["expiringOn"] = d.ExpiringOn.Time.Format(time.RFC3339)
	}
	return payload
}

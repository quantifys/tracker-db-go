package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	NotificationType int32          `gorm:"type:integer;column:alert_type" json:"alertType"`
	IsRead           bool           `gorm:"column:is_read;default:false" json:"isRead"`
	Attributes       datatypes.JSON `gorm:"column:attributes" json:"attributes"`
	UserId           uuid.UUID      `gorm:"index:,sort:desc;column:user_id" json:"userId"`
	User             User           `gorm:"foreignKey:UserId"`
}

func (Notification) TableName() string {
	return "alerts"
}

func (d Notification) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         d.ID,
		"alertType":  d.NotificationType,
		"userId":     d.UserId,
		"attributes": d.Attributes,
		"isRead":     d.IsRead,
		"createdAt":  d.CreatedAt.Format(time.RFC3339),
	}
	return payload
}

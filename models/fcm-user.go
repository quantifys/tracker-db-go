package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FcmUser struct {
	gorm.Model
	Id            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId        *uuid.UUID `gorm:"type:uuid;column:user_id" json:"userId"`
	User          User       `gorm:"foreignKey:UserId"`
	FcmToken      string     `gorm:"column:fcm_token;" json:"fcmToken"`
	Device        string     `gorm:"column:device;" json:"device"`
	DeviceId      string     `gorm:"column:device_id;" json:"deviceId"`
	Platform      int16      `gorm:"column:platform;" json:"platform"`
	IsInvalidated bool       `gorm:"column:is_invalidated;" json:"isInvalidated"`
	AppType       int16      `gorm:"column:app_type;" json:"appType"`
}

func (FcmUser) TableName() string {
	return "fcm_users"
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	SerialNo string    `gorm:"column:serial_no;unique" json:"serialNo"`
	Imei1    string    `gorm:"column:imei_1;unique" json:"imei1"`
	Imei2    string    `gorm:"column:imei_2;unique" json:"imei2"`
	Phone1   string    `gorm:"column:phone_1;unique" json:"phone1"`
	Phone2   string    `gorm:"column:phone_2;unique" json:"phone2"`
	Type     int16     `gorm:"column:type;" json:"type"`
}

func (Device) TableName() string {
	return "devices"
}

func (u Device) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"serialNo":  u.SerialNo,
		"imei1":     u.Imei1,
		"imei2":     u.Imei2,
		"phone1":    u.Phone1,
		"phone2":    u.Phone2,
		"type":      u.Type,
		"createdAt": u.CreatedAt,
	}
	return payload
}

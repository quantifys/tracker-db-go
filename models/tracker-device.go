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
	Type     string    `gorm:"column:type;" json:"type"`
}

func (Device) TableName() string {
	return "devices"
}
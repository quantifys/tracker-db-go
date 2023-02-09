package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Id       uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	SerialNo string         `gorm:"column:serial_no;unique" json:"serialNo"`
	Imei     string         `gorm:"column:imei;unique" json:"imei"`
	Esim1    string         `gorm:"column:esim_1;unique" json:"esim1"`
	Esim2    sql.NullString `gorm:"column:esim_2;unique;default:null" json:"esim2"`
	Phone1   string         `gorm:"column:phone_1;unique" json:"phone1"`
	Phone2   sql.NullString `gorm:"column:phone_2;unique;default:null" json:"phone2"`
	Iccid    sql.NullString `gorm:"column:iccid;unique;default:null" json:"iccid"`
	Type     int16          `gorm:"column:type;" json:"type"`
}

func (Device) TableName() string {
	return "devices"
}

func (u Device) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":        u.Id,
		"serialNo":  u.SerialNo,
		"imei":      u.Imei,
		"esim1":     u.Esim1,
		"esim2":     u.Esim2,
		"phone1":    u.Phone1,
		"phone2":    u.Phone2,
		"iccid":     u.Iccid,
		"type":      u.Type,
		"createdAt": u.CreatedAt,
	}
	return payload
}

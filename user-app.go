package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserApp struct {
	gorm.Model
	Id        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId    *uuid.UUID `gorm:"type:uuid;column:user_id" json:"userId"`
	User      User       `gorm:"foreignKey:UserId"`
	Name      string     `gorm:"column:name;" json:"name"`
	AccessKey string     `gorm:"column:access_key;" json:"accessKey"`
	Active    bool       `gorm:"column:active;" json:"active"`
}

func (UserApp) TableName() string {
	return "user_app"
}

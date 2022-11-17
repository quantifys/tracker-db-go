package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name         string         `gorm:"column:name" json:"name"`
	Phone        string         `gorm:"column:phone;unique" json:"phone"`
	Email        sql.NullString `gorm:"column:email;unique" json:"email"`
	Username     sql.NullString `gorm:"column:username;unique" json:"username"`
	Role         string         `gorm:"column:role;default:end_user" json:"role"`
	PasswordHash sql.NullString `gorm:"column:password_hash" json:"passwordHash"`
	ProfilePic   sql.NullString `gorm:"column:profile_pic" json:"profilePic"`
}

func (User) TableName() string {
	return "users"
}

func (u User) ShortJson() map[string]interface{} {
	payload := map[string]interface{}{
		"id":    u.Id,
		"name":  u.Name,
		"email": u.Email,
		"phone": u.Phone,
		"role":  u.Role,
	}
	return payload
}

func (u User) Json() map[string]interface{} {
	payload := map[string]interface{}{
		"id":         u.Id,
		"name":       u.Name,
		"email":      nil,
		"phone":      u.Phone,
		"role":       u.Role,
		"profilePic": nil,
		"gender":     nil,
		"birthDate":  nil,
		"createdAt":  u.CreatedAt.Format(time.RFC3339),
	}
	if u.Email.Valid {
		payload["email"] = u.Email.String
	}
	if u.ProfilePic.Valid {
		payload["profilePic"] = u.ProfilePic.String
	}
	return payload
}

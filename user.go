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
	Username     string         `gorm:"column:username;unique" json:"username"`
	Role         int16          `gorm:"column:role;default:3" json:"role"`
	PasswordHash sql.NullString `gorm:"column:password_hash" json:"passwordHash"`
	ProfilePic   sql.NullString `gorm:"column:profile_pic" json:"profilePic"`
	Gender       sql.NullInt16  `gorm:"column:gender" json:"gender"`
	BirthDate    sql.NullTime   `gorm:"column:birth_date" json:"birthDate"`
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
	if u.Gender.Valid {
		payload["gender"] = u.Gender.Int16
	}
	if u.BirthDate.Valid {
		payload["birthDate"] = u.BirthDate.Time.Format(time.RFC3339)
	}
	return payload
}

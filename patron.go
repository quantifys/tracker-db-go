package models

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Patron struct {
	gorm.Model
	Id           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId       *uuid.UUID     `gorm:"type:uuid;column:user_id" json:"userId"`
	User         User           `gorm:"foreignKey:UserId"`
	VendorId     *uuid.UUID     `gorm:"type:uuid;column:vendor_id" json:"vendorId"`
	Vendor       User           `gorm:"foreignKey:VendorId"`
	TrackingTags pq.StringArray `gorm:"column:tracking_tags;" json:"trackingTags"`
	Address1     string         `gorm:"column:address_1;" json:"address1"`
	Address2     sql.NullString `gorm:"column:address_2;" json:"address2"`
	Locality     string         `gorm:"column:locality;" json:"locality"`
	City         string         `gorm:"column:city;" json:"city"`
	PostalCode   string         `gorm:"column:postal_code;" json:"postalCode"`
	State        string         `gorm:"column:state;" json:"state"`
	Country      string         `gorm:"column:country;" json:"country"`
	Gstn         sql.NullString `gorm:"column:gstn;unique" json:"gstn"`
	Active       bool           `gorm:"column:active;default:true" json:"active"`
}

func (Patron) TableName() string {
	return "patron"
}

func (u Patron) Json(preloaded bool) map[string]interface{} {
	payload := map[string]interface{}{
		"id":           u.Id,
		"trackingTags": u.TrackingTags,
		"address1":     u.Address1,
		"address2":     u.Address2,
		"locality":     u.Locality,
		"city":         u.City,
		"postalCode":   u.PostalCode,
		"state":        u.State,
		"country":      u.Country,
		"gstn":         u.Gstn,
		"active":       u.Active,
	}
	if preloaded {
		payload["user"] = u.User.ShortJson()
		payload["vendor"] = u.Vendor.ShortJson()
	} else {
		payload["userId"] = u.UserId
		payload["vendorId"] = u.VendorId
	}
	return payload
}

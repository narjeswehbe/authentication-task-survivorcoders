package entity

import "time"

type BlackList struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint64    `gorm:";not null;" json:"user_id"`
	AuthUUID   string    `gorm:"size:255;not null;" json:"auth_uuid"`
	ExpiryDate time.Time `gorm:"size:255;not null;" json:"expiry_date"`
}

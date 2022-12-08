package entity

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Expires_at int64
	UserId     uint
	Token      string
}

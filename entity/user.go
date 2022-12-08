package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Email             string `gorm:"unique""`
	Username          string `gorm:"unique""`
	Verification_code string
	Verified          bool
	//optional fields
	Phone_number string
	Nationality  string
	Password     string
}

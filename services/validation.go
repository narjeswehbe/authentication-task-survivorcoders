package services

import (
	"net/mail"
	"unicode"
)

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validPassword(password string) bool {
	var (
		upp, low, num bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true

		case unicode.IsLower(char):
			low = true

		case unicode.IsNumber(char):
			num = true
		default:
			return false
		}
	}

	if !upp || !low || !num || len(password) < 8 {
		return false
	}

	return true
}

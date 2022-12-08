package services

import (
	"github.com/matthewhartstonge/argon2"
)

func HashPassword(password string) string {
	argon := argon2.DefaultConfig()

	// Waaahht??! It includes magic salt generation for me ! Yasss...
	encoded, err := argon.HashEncoded([]byte(password))
	if err != nil {
		panic(err) // 💥
	}

	return string(encoded)

}

func VerifyPassword(pass string, hash string) bool {

	ok, err := argon2.VerifyEncoded([]byte(pass), []byte(hash))
	if err != nil {
		panic(err)

	}
	return ok

}

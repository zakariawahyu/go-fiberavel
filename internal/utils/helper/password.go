package helper

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Printf("Failed to generate password: %v", err)
		return ""
	}

	return string(hashed)
}

func ComparePassword(hashPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
}

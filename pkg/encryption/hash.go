package encryption

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	saltString := hex.EncodeToString(salt)

	return saltString, nil
}

func HashPassword(password string, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(hashed, password, salt string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password+salt))

	return err
}
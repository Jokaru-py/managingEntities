package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func PasswordCheck(dbPass string, tekPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(tekPass))
}

func EndcodeLink(str string) string {
	var result string
	for _, value := range str {
		result += string(value - 1)
	}

	return result
}

func DecodeLink(str string) string {
	var result string
	for _, value := range str {
		result += string(value + 1)
	}

	return result
}

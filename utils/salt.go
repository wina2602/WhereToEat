package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func Createsalt() (string, error) {
	length := 16
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

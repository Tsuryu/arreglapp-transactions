package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// RandomHex : generates a random hex value
func RandomHex(lenght int) (string, error) {
	bytes := make([]byte, lenght)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(bytes)), nil
}

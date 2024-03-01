package pwhash

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

// Hash hashes password with given salt
func Hash(salt string, password string) (string, error) {
	hash, err := scrypt.Key(
		[]byte(password),
		[]byte(salt),
		32768,
		8,
		1,
		32)
	if err != nil {
		return "", fmt.Errorf("scrypt failed: %w", err)
	}

	return hex.EncodeToString(hash), nil
}

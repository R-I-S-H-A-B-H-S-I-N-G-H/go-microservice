package hashing_util

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// The length is in characters, so it should be an even number (for hex encoding).
func GenerateRandomHash(length int) (string, error) {
	// Since each byte is represented by 2 hex characters, length must be even
	if length%2 != 0 {
		return "", errors.New("length must be an even number")
	}

	// Create a byte slice of half the size (since each byte = 2 hex characters)
	bytesLength := length / 2
	bytes := make([]byte, bytesLength)

	// Generate secure random bytes
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a hex string
	randomHash := hex.EncodeToString(bytes)
	return randomHash, nil
}

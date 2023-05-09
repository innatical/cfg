package main

import (
	"crypto/rand"
	"encoding/base64"
)

func generateSecureString(length int) (string, error) {
	// Since base64 encoding increases the size by 4/3, we need to adjust the number of random bytes
	randomBytes := make([]byte, (length*3+3)/4)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a base64-encoded string
	encoded := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Truncate the encoded string to the desired length
	return encoded[:length], nil
}

package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"golang.org/x/crypto/argon2"
)

const (
	saltLength = 16
	timeCost   = 1
	memoryCost = 64 * 1024
	threads    = 4
	keyLength  = 32
)

func GenerateSalt() (string, error) {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func Hash(password, salt string) string {
	hash := argon2.IDKey([]byte(password), []byte(salt), timeCost, memoryCost, threads, keyLength)
	return hex.EncodeToString(hash)
}

func VerifyPassword(storedHash, salt, password string) bool {
	// Hash the provided password with the stored salt
	hash := argon2.IDKey([]byte(password), []byte(salt), timeCost, memoryCost, threads, keyLength)
	hashedPassword := hex.EncodeToString(hash)

	// Use constant-time comparison to prevent timing attacks
	return subtle.ConstantTimeCompare([]byte(storedHash), []byte(hashedPassword)) == 1
}

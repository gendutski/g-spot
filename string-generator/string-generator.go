package stringgenerator

import (
	cryptoRand "crypto/rand"
	"encoding/base64"
	mathRand "math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generate token
func GenerateToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// generate random alphanumeric string
func Serialize(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[mathRand.Intn(len(letterBytes))]
	}
	return string(b)
}

package helper

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/a-fandy/finan/exception"
	"golang.org/x/crypto/bcrypt"
)

func ConvertStringToUint64(num string) uint64 {
	// Convert string to uint64
	uint64Value, err := strconv.ParseUint(num, 10, 64)
	if err != nil {
		exception.PanicIfError(err)
	}
	return uint64Value
}

func HashingPassword(password string) string {
	passwordByte := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordByte, 15)
	exception.PanicIfError(err)
	return string(hash)
}

func CheckPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Define the characters that can be used in the random string
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate the random string
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}

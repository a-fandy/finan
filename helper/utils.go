package helper

import (
	"strconv"

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

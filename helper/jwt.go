package helper

import (
	"crypto/rsa"
	"time"

	userEntity "github.com/a-fandy/finan/api/v1/user"
	"github.com/a-fandy/finan/exception"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(user userEntity.User, secretKey *rsa.PrivateKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":  "finan-api",
		"sub":  user.Email,
		"role": user.Role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().AddDate(0, 1, 0).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	exception.PanicIfError(err)
	return tokenString
}

func VerifyJWTToken(tokenString string, secretKey *rsa.PublicKey) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, exception.UnauthorizedError{Message: "Unexpected signing method"}
		}
		return secretKey, nil
	})
	if err != nil {
		panic(exception.UnauthorizedError{Message: err.Error()})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic(exception.UnauthorizedError{Message: "Failed parse token"})
	}
	return claims
}

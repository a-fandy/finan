package middleware

import (
	"strings"

	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	config.Config
}

func NewAuthMiddleware(config config.Config) *AuthMiddleware {
	return &AuthMiddleware{Config: config}
}

func (authMiddleware AuthMiddleware) AuthenticatedJWT(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	claims := helper.VerifyJWTToken(token, authMiddleware.GetPublicKey())
	if claims["role"] != "user" {
		panic(exception.ForbiddenError{Message: "Forbidden Access"})
	}
	ctx.Locals("user", claims["sub"])
	return ctx.Next()
}

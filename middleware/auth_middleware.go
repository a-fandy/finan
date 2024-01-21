package middleware

import (
	"github.com/a-fandy/finan/exception"
	"github.com/gofiber/fiber/v2"
)

func AuthenticatedJWT(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	// Bearer
	if token != "Bearer secret" {
		panic(exception.UnauthorizedError{Message: "Unauthorized"})
	}
	return ctx.Next()
}

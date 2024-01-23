package login

import (
	"github.com/a-fandy/finan/api/v1/user/repository"
	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/gofiber/fiber/v2"
)

type Authetication struct {
	repository.UserRepository
	config.Config
}

func NewAuthentication(userRepository *repository.UserRepository, config config.Config) *Authetication {
	return &Authetication{UserRepository: *userRepository, Config: config}
}

func (authetication Authetication) Login(ctx *fiber.Ctx) error {
	var request LoginRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)

	helper.Validate(request)

	user, err := authetication.UserRepository.FindByEmail(ctx.Context(), request.Email)
	if err != nil {
		panic(exception.UnauthorizedError{Message: "Authentication failed. Please check your credentials and try again."})
	}

	if !helper.CheckPasswordHash(user.Password, request.Password) {
		panic(exception.UnauthorizedError{Message: "Authentication failed. Please check your credentials and try again."})
	}

	return ctx.Status(fiber.StatusOK).JSON(helper.NewSuccessResponse(AuthToLoginResponse(user, authetication.Config)))
}

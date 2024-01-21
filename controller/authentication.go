package controller

import (
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/a-fandy/finan/model/web"
	"github.com/a-fandy/finan/repository"
	"github.com/gofiber/fiber/v2"
)

type Authetication struct {
	repository.UserRepository
}

func NewAuthentication(userRepository *repository.UserRepository) *Authetication {
	return &Authetication{UserRepository: *userRepository}
}

func (authetication Authetication) Login(ctx *fiber.Ctx) error {
	var request web.LoginRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)

	helper.Validate(request)

	user, err := authetication.UserRepository.FindByEmail(ctx.Context(), request.Email)
	if err != nil {
		panic(exception.UnauthorizedError{Message: "Authentication failed. Please check your credentials and try again."})
	}

	if helper.CheckPasswordHash(user.Password, request.Password) {
		return ctx.Status(fiber.StatusOK).JSON(web.NewSuccessResponse(helper.AuthToLoginResponse(user)))
	}

	panic(exception.UnauthorizedError{Message: "Authentication failed. Please check your credentials and try again."})
}

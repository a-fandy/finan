package register

import (
	userEntity "github.com/a-fandy/finan/api/v1/user"
	"github.com/a-fandy/finan/api/v1/user/repository"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/gofiber/fiber/v2"
)

type RegisterController struct {
	repository.UserRepository
}

func NewRegister(userRepository *repository.UserRepository) *RegisterController {
	return &RegisterController{UserRepository: *userRepository}
}

func (register RegisterController) Register(ctx *fiber.Ctx) error {
	var request RegisterRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)
	helper.Validate(request)

	user := RegisterRequestToEntity(request)
	user = register.UserRepository.Insert(ctx.Context(), user)

	response := userEntity.UserEntityToResponse(user)
	return ctx.Status(fiber.StatusCreated).JSON(helper.NewSuccessResponse(response))
}

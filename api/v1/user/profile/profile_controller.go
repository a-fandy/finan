package profile

import (
	userEntity "github.com/a-fandy/finan/api/v1/user"
	"github.com/a-fandy/finan/api/v1/user/repository"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	repository.UserRepository
}

func NewProfile(userRepository *repository.UserRepository) *ProfileController {
	return &ProfileController{UserRepository: *userRepository}
}

func (profile ProfileController) Profile(ctx *fiber.Ctx) error {
	id := ctx.Locals("user")
	user, err := profile.UserRepository.FindByEmail(ctx.Context(), id.(string))
	exception.PanicIfError(err)
	result := userEntity.UserEntityToResponse(user)
	return ctx.Status(fiber.StatusOK).JSON(helper.NewSuccessResponse(result))
}

func (profile ProfileController) ProfileUpdate(ctx *fiber.Ctx) error {
	var request UserRequestUpdate
	id := ctx.Locals("user")
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)

	helper.Validate(request)
	userUpdate := UserRequestUpdateToEntity(request)
	user, err := profile.UserRepository.FindByEmail(ctx.Context(), id.(string))
	exception.PanicIfError(err)
	user = profile.UserRepository.Update(ctx.Context(), user, userUpdate)
	response := userEntity.UserEntityToResponse(user)
	return ctx.Status(fiber.StatusOK).JSON(helper.NewSuccessResponse(response))
}

func (profile ProfileController) ProfileNonactive(ctx *fiber.Ctx) error {
	id := ctx.Locals("user")
	user, err := profile.UserRepository.FindByEmail(ctx.Context(), id.(string))
	exception.PanicIfError(err)
	profile.UserRepository.Delete(ctx.Context(), user)
	return ctx.Status(fiber.StatusOK).JSON(helper.NewSuccessResponse(nil))
}

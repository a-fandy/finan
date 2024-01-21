package controller

import (
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/a-fandy/finan/model/web"
	"github.com/a-fandy/finan/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: *userService}
}

func (controller UserController) Create(ctx *fiber.Ctx) error {

	var request web.UserRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)

	response := controller.UserService.Create(ctx.Context(), request)
	return ctx.Status(fiber.StatusCreated).JSON(web.NewSuccessResponse(response))
}

func (controller UserController) Update(ctx *fiber.Ctx) error {
	var request web.UserRequestUpdate
	id := ctx.Locals("user")
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)

	response := controller.UserService.Update(ctx.Context(), request, id.(string))
	return ctx.Status(fiber.StatusOK).JSON(web.NewSuccessResponse(response))
}

func (controller UserController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	controller.UserService.Delete(ctx.Context(), helper.ConvertStringToUint64(id))
	return ctx.Status(fiber.StatusOK).JSON(web.NewSuccessResponse(nil))
}

func (controller UserController) FindById(ctx *fiber.Ctx) error {
	id := ctx.Locals("user")
	result := controller.UserService.FindById(ctx.Context(), id.(string))
	return ctx.Status(fiber.StatusOK).JSON(web.NewSuccessResponse(result))
}

func (controller UserController) FindAll(ctx *fiber.Ctx) error {
	result := controller.UserService.FindAll(ctx.Context())
	return ctx.Status(fiber.StatusOK).JSON(web.NewSuccessResponse(result))
}

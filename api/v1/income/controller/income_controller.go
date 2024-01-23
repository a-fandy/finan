package income

import (
	"fmt"

	incomeEntity "github.com/a-fandy/finan/api/v1/income"
	"github.com/a-fandy/finan/api/v1/income/repository"
	userRepo "github.com/a-fandy/finan/api/v1/user/repository"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/helper"
	"github.com/gofiber/fiber/v2"
)

type IncomeController struct {
	repository.IncomeRepository
	userRepo.UserRepository
}

func NewIncomeController(incomeRepository *repository.IncomeRepository, userRepository *userRepo.UserRepository) *IncomeController {
	return &IncomeController{IncomeRepository: *incomeRepository, UserRepository: *userRepository}
}

func (controller IncomeController) Create(ctx *fiber.Ctx) error {
	var request incomeEntity.IncomeRequest
	id := ctx.Locals("user")
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)
	helper.Validate(request)
	income := incomeEntity.IncomeRequestToEntity(request)
	user, err := controller.UserRepository.FindByEmail(ctx.Context(), id.(string))
	exception.PanicIfError(err)
	income = controller.IncomeRepository.Insert(ctx.Context(), user, income)

	response := incomeEntity.IncomeEntityToResponse(income)
	fmt.Println(response, income)
	return ctx.Status(fiber.StatusCreated).JSON(helper.NewSuccessResponse(response))
}

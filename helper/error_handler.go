package helper

import (
	"encoding/json"

	"github.com/a-fandy/finan/exception"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type ErrHandler struct {
	*mongo.Client
}

func NewErrorHandler(mongoClient *mongo.Client) *ErrHandler {
	return &ErrHandler{mongoClient}
}

func (errHandler ErrHandler) ErrorHandler(ctx *fiber.Ctx, err error) error {
	LogErrorToMongoDB(ctx, err, errHandler.Client)

	_, validationError := err.(exception.ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		exception.PanicIfError(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("Bad Request", messages))
	}

	_, notFoundError := err.(exception.NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(NewErrorResponse("Not Found", err.Error()))
	}

	_, unauthorizedError := err.(exception.UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(NewErrorResponse("Unauthorized", err.Error()))
	}

	_, forbiddenError := err.(exception.ForbiddenError)
	if forbiddenError {
		return ctx.Status(fiber.StatusForbidden).JSON(NewErrorResponse("Forbidden", err.Error()))
	}

	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		// Check if the error is a duplicate entry error (Error 1062)
		if mysqlErr.Number == 1062 {
			return ctx.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("Bad Request", mysqlErr.Message))
		}
	}
	log.Error(err.Error())
	return ctx.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse("General Error", err.Error()))
}

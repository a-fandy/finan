package exception

import (
	"encoding/json"

	"github.com/a-fandy/finan/model/web"
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

	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicIfError(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(web.NewErrorResponse("Bad Request", messages))
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(web.NewErrorResponse("Not Found", err.Error()))
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.NewErrorResponse("Unauthorized", err.Error()))
	}

	_, forbiddenError := err.(ForbiddenError)
	if forbiddenError {
		return ctx.Status(fiber.StatusForbidden).JSON(web.NewErrorResponse("Forbidden", err.Error()))
	}

	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		// Check if the error is a duplicate entry error (Error 1062)
		if mysqlErr.Number == 1062 {
			return ctx.Status(fiber.StatusBadRequest).JSON(web.NewErrorResponse("Bad Request", mysqlErr.Message))
		}
	}
	log.Error(err.Error())
	return ctx.Status(fiber.StatusInternalServerError).JSON(web.NewErrorResponse("General Error", err.Error()))
}

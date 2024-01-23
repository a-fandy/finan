package config

import (
	"github.com/a-fandy/finan/helper"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewFiberConfiguration(mongoClient *mongo.Client) fiber.Config {
	handler := helper.NewErrorHandler(mongoClient)
	return fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	}
}

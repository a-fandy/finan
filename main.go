package main

import (
	"os"
	"time"

	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	//setup config
	configuration := config.New()

	// Initial database
	database := config.ConnectDatabase(configuration)
	mongoClient, err := config.ConnectMongo(configuration)
	if err != nil {
		log.Error("Error initializing MongoDB")
		os.Exit(1)
	}

	//setup fiber
	app := fiber.New(config.NewFiberConfiguration(mongoClient))
	// Middleware to log requests
	app.Use(recover.New())
	// app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("start_time", time.Now())
		if err := c.Next(); err != nil {
			// config.LogErrorToMongoDB(c.Context(), err, c, mongoClient)
			return err
		}
		return nil
	})

	//routing
	route.RouteInit(app, database, configuration)

	//start app
	err = app.Listen(":3000")
	exception.PanicIfError(err)
}

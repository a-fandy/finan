package main

import (
	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/exception"
	"github.com/a-fandy/finan/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	//setup config
	configuration := config.New()
	// Initial database
	database := config.ConnectDatabase(configuration)

	//setup fiber
	app := fiber.New(config.NewFiberConfiguration())
	app.Use(recover.New())

	//routing
	route.RouteInit(app, database, configuration)

	//start app
	err := app.Listen(":3000")
	exception.PanicIfError(err)
}

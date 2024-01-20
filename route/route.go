package route

import (
	"github.com/a-fandy/finan/controller"
	repository "github.com/a-fandy/finan/repository/impl"
	service "github.com/a-fandy/finan/service/impl"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteInit(r *fiber.App, DB *gorm.DB) {
	api := r.Group("/api")

	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)
	api.Post("/user", userController.Create)
	api.Put("/user/:id", userController.Update)
	api.Delete("/user/:id", userController.Delete)
	api.Get("/user/:id", userController.FindById)
	api.Get("/user", userController.FindAll)

}

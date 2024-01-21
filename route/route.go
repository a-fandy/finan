package route

import (
	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/controller"
	"github.com/a-fandy/finan/middleware"
	repository "github.com/a-fandy/finan/repository/impl"
	service "github.com/a-fandy/finan/service/impl"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteInit(r *fiber.App, DB *gorm.DB, config config.Config) {
	api := r.Group("/api")
	authMiddleware := middleware.NewAuthMiddleware(config)

	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(&userRepository)
	userController := controller.NewUserController(&userService)

	// api.Delete("/user/:id", userController.Delete)
	// api.Get("/user", authMiddleware.AuthenticatedJWT, userController.FindAll)

	auth := api.Group("/auth")
	authController := controller.NewAuthentication(&userRepository, config)
	auth.Post("/login", authController.Login)
	auth.Post("/register", userController.Create)
	auth.Get("/", authMiddleware.AuthenticatedJWT, userController.FindById)
	auth.Put("/", authMiddleware.AuthenticatedJWT, userController.Update)

}

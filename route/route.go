package route

import (
	incomeCon "github.com/a-fandy/finan/api/v1/income/controller"
	incomeRepo "github.com/a-fandy/finan/api/v1/income/repository"
	"github.com/a-fandy/finan/api/v1/user/login"
	"github.com/a-fandy/finan/api/v1/user/profile"
	"github.com/a-fandy/finan/api/v1/user/register"
	userRepo "github.com/a-fandy/finan/api/v1/user/repository"
	"github.com/a-fandy/finan/config"
	"github.com/a-fandy/finan/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteInit(r *fiber.App, DB *gorm.DB, config config.Config) {
	api := r.Group("/api")
	authMiddleware := middleware.NewAuthMiddleware(config)

	userRepository := userRepo.NewUserRepository(DB)
	loginController := login.NewAuthentication(&userRepository, config)
	registerController := register.NewRegister(&userRepository)
	profileController := profile.NewProfile(&userRepository)
	// api.Delete("/user/:id", userController.Delete)
	// api.Get("/user", authMiddleware.AuthenticatedJWT, userController.FindAll)

	auth := api.Group("/auth")
	auth.Post("/login", loginController.Login)
	auth.Post("/register", registerController.Register)
	auth.Get("/", authMiddleware.AuthenticatedJWT, profileController.Profile)
	auth.Put("/", authMiddleware.AuthenticatedJWT, profileController.ProfileUpdate)

	incomeRepository := incomeRepo.NewIncomeRepository(DB)
	incomeController := incomeCon.NewIncomeController(&incomeRepository, &userRepository)
	api.Post("/income", authMiddleware.AuthenticatedJWT, incomeController.Create)

}

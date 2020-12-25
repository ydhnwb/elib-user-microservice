package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-user-microservice/application/middleware"
	"github.com/ydhnwb/elib-user-microservice/application/repository"
	"github.com/ydhnwb/elib-user-microservice/application/service/authservice"
	"github.com/ydhnwb/elib-user-microservice/application/service/userservice"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/controller"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/persistence"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = persistence.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	authService    authservice.AuthService   = authservice.NewAuthService(userRepository)
	jwtService     authservice.JWTService    = authservice.NewJWTService()
	userService    userservice.UserService   = userservice.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(jwtService, authService)
	userController controller.UserController = controller.NewUserController(jwtService, userService)
)

func main() {
	defer persistence.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run(":8080")

}

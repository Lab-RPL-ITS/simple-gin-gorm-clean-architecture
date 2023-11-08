package main

import (
	"rpl-simple-backend/config"
	"rpl-simple-backend/controller"
	"rpl-simple-backend/middleware"
	"rpl-simple-backend/repository"
	"rpl-simple-backend/routes"
	"rpl-simple-backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db *gorm.DB = config.SetUpDatabaseConnection()

		userRepo              repository.UserRepository        = repository.NewUserRepository(db)
		phoneNumberRepo       repository.PhoneNumberRepository = repository.NewPhoneNumberRepository(db)

		userService           service.UserService              = service.NewUserService(userRepo, phoneNumberRepo)
		phoneNumberService    service.PhoneNumberService       = service.NewPhoneNumberService(phoneNumberRepo, userRepo)

		userController        controller.UserController        = controller.NewUserController(userService)
		phoneNumberController controller.PhoneNumberController = controller.NewPhoneNumberController(phoneNumberService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.User(server, userController)
	routes.PhoneNumber(server, phoneNumberController)

	port := "8080"
	server.Run(":" + port)
}

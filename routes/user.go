package routes

import (
	"rpl-simple-backend/controller"

	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController) {
	routes := route.Group("/api/user") 
	{
		routes.POST("", userController.CreateUser)
		routes.GET("", userController.GetAllUser)
		routes.GET("/:id", userController.GetUserById)
	}
}
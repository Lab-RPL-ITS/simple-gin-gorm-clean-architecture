package routes

import (
	"rpl-simple-backend/controller"

	"github.com/gin-gonic/gin"
)

func PhoneNumber(route *gin.Engine, phoneNumber controller.PhoneNumberController) {
	routes := route.Group("/api/phone-number")
	{
		routes.POST("", phoneNumber.CreatePhoneNumber)
		routes.GET("", phoneNumber.GetAllPhoneNumber)
		routes.GET("/:id", phoneNumber.GetPhoneNumberById)
	}
}
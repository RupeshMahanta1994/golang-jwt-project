package routes

import (
	controller "github.com/RupeshMahanta1994/go-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/singup", controller.SignUp)
	// incomingRoutes.POST("users/login", controller.Login())

}

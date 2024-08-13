package routes

import (
	"github.com/dilippm92/taskmanager/controllers"
	"github.com/gin-gonic/gin"
	"github.com/dilippm92/taskmanager/middlewares"

)

func Authroutes(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/auth")
	{
		authGroup.GET("/test", controllers.TestController) // Handle POST /auth/login
		authGroup.GET("/get_user/:id", middlewares.JwtTokenVerify(), controllers.GetUserByID)
		authGroup.POST("/user_signup", controllers.CreateNewUser)
		authGroup.POST("/user_login", controllers.UserLogin)
	}

}

package routes

import (
	"github.com/dilippm92/taskmanager/controllers"
	"github.com/gin-gonic/gin"
)

func Authroutes(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/auth")
	{
		authGroup.GET("/test", controllers.TestController) // Handle POST /auth/login
		authGroup.GET("/get_user/:id", controllers.GetUserByID)
		authGroup.POST("/user_signup", controllers.CreateNewUser)
		authGroup.POST("/user_login", controllers.UserLogin)
	}

}

package routes

import (
	"github.com/dilippm92/taskmanager/controllers"
	"github.com/dilippm92/taskmanager/middlewares"
	"github.com/gin-gonic/gin"
)
func Sprintroutes(routerGroup *gin.RouterGroup)  {
	sprintGroup:=routerGroup.Group("/sprint")
	{
		sprintGroup.GET("/test",middlewares.JwtTokenVerify(),controllers.TestSprintController)
		sprintGroup.POST("/create_sprint",middlewares.JwtTokenVerify(),controllers.CreateNewSprint)
	}
}
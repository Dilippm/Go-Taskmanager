package routes

import (
	"github.com/dilippm92/taskmanager/controllers"
	"github.com/dilippm92/taskmanager/middlewares"
	"github.com/gin-gonic/gin"
)
func Sprintroutes(routerGroup *gin.RouterGroup)  {
	sprintGroup:=routerGroup.Group("/sprint",middlewares.JwtTokenVerify())
	{
		sprintGroup.GET("/test",controllers.TestSprintController)
		sprintGroup.POST("/create_sprint",controllers.CreateNewSprint)
		sprintGroup.GET("/get_all_sprints/:id",controllers.GetAllSprints)
		sprintGroup.GET("/get_sprint_detail/:id",controllers.GetSprintDetails)
		sprintGroup.PUT("/update_sprint/:id",controllers.UpdateSprint)
		sprintGroup.DELETE("/delete_sprint/:id",controllers.DeleteSprint)
	}
}
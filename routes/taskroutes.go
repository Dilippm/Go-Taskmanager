package routes

import (
	"github.com/dilippm92/taskmanager/controllers"
	"github.com/dilippm92/taskmanager/middlewares"
	"github.com/gin-gonic/gin"
)
func Taskroutes(routerGroup *gin.RouterGroup){
	taskGroup:=routerGroup.Group("/task",middlewares.JwtTokenVerify())
{
	taskGroup.GET("/test",controllers.TestTaskController)
	taskGroup.POST("/create_task",controllers.CreateNewTask)
	taskGroup.GET("/get_task_details/:id",controllers.GetTaskDetails)
	taskGroup.PUT("/update_task/:id",controllers.UpdateTask)
	taskGroup.DELETE("/delete_task/:id",controllers.DeleteTask)
}
}
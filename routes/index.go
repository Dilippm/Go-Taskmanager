package routes

import "github.com/gin-gonic/gin"

func MainRoutes(router *gin.Engine) {
	// Define /api base path
	apiGroup := router.Group("/api")

	Authroutes(apiGroup)
	Sprintroutes(apiGroup)
	Taskroutes(apiGroup)
}

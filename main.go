package main

import (
	"github.com/dilippm92/taskmanager/config"
	"github.com/dilippm92/taskmanager/middlewares"
	"github.com/dilippm92/taskmanager/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := config.PORT
	config.ConnectMongoDB()
	if port == "" {

		port = "8000"

	}
	router := gin.New()
	router.Use(middlewares.ErrorHandler()) // Use the middleware correctly
	router.Use(gin.Logger())
	routes.MainRoutes(router)
	router.Run(":" + port)
}

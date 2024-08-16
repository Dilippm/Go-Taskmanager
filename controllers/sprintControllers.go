package controllers

import (
	"fmt"
	"net/http"

	"github.com/dilippm92/taskmanager/models"
	"github.com/dilippm92/taskmanager/models/queries"
	"github.com/gin-gonic/gin"
)

// test controller for sprint controllers
func TestSprintController(c *gin.Context){
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// Use the userId (make sure to assert it to the appropriate type)
	userIdStr := userId.(string)

	c.JSON(http.StatusOK,gin.H{"message": fmt.Sprintf("sprint test route: %v", userIdStr),})
}

// create a new sprint
func CreateNewSprint(c *gin.Context){
	var sprint models.Sprint
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// Use the userId (make sure to assert it to the appropriate type)
	userIdStr := userId.(string)
	// bind the request body to the Sprint struct
	if err:= c.ShouldBindJSON(&sprint);err!=nil{
		c.Error(fmt.Errorf("failed to bind request body to sprint model:%v",err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// Set the UserId field in the Sprint struct
	sprint.UserId = userIdStr
	result,err:= queries.CreateSprint(sprint)
	if err!=nil{
		c.Error(fmt.Errorf("failed to create sprint in the database: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// return a sucess response with the result
	c.JSON(http.StatusCreated,gin.H{
		"message":"Sprint created successfully",
		"result":result.InsertedID,
	})
}
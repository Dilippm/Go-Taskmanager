package controllers

import (
	"fmt"
	"net/http"

	"github.com/dilippm92/taskmanager/models"
	"github.com/dilippm92/taskmanager/models/queries"
	"github.com/gin-gonic/gin"
)

// test controller for task controllers
func TestTaskController(c *gin.Context){
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// Use the userId (make sure to assert it to the appropriate type)
	userIdStr := userId.(string)

	c.JSON(http.StatusOK,gin.H{"message": fmt.Sprintf("sprint test route: %v", userIdStr),})
}

// function to create a new task
func CreateNewTask(c *gin.Context){
	var task models.SubTask
	_,exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	// bind the body to task struct
	if err:= c.ShouldBindJSON(&task);err!=nil{
		c.Error(fmt.Errorf("failed to bind request body to task model:%v",err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result,err:= queries.CreateSubTask(task)
	if err!=nil{
		c.Error(fmt.Errorf("failed to create sprint in the database: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
// return a sucess response with the result
c.JSON(http.StatusCreated,gin.H{
	"message":"Task created successfully",
	"result":result.InsertedID,
})
}
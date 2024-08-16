package controllers

import (
	"fmt"
	"net/http"

	"github.com/dilippm92/taskmanager/models"
	"github.com/dilippm92/taskmanager/models/queries"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// get all sprints fo a user by userId
func GetAllSprints(c *gin.Context){
	userId := c.Param("id")
	
	sprints,err:= queries.GetAllSprints(userId)
	if err!= nil{
		c.Error(fmt.Errorf("failed to find sprints with userID %s: %w", userId, err))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, sprints)
}

// get sprint detials by sprint id
func GetSprintDetails(c *gin.Context){
	sprintId:=c.Param("id")
	sprint,err:=queries.GetSingleSprintDetails(sprintId)
	if err!= nil{
		c.Error(fmt.Errorf("failed to find sprints with sprintID %s: %w", sprintId, err))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK,sprint)
}

// update a sprint by sprint id
func UpdateSprint(c *gin.Context){
	id:= c.Param("id")
	sprintID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sprint ID"})
        return
    }
	 // Bind the request body to the Sprint struct
	 var sprint models.Sprint
	 if err := c.ShouldBindJSON(&sprint); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		 return
	 }
 
	 // Update the sprint in the database
	 result, err := queries.UpdateSprint(sprintID, sprint)
	 if err != nil {
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sprint"})
		 return
	 }
 
	 // Return a success response
	 c.JSON(http.StatusOK, gin.H{
		 "message":    "Sprint updated successfully",
		 "matched":    result.MatchedCount,
		 "modified":   result.ModifiedCount,
		 "upsertedID": result.UpsertedID,
	 })

}

// delete a sprint by sprint id
func DeleteSprint(c *gin.Context){
	id:= c.Param("id")
	sprintID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sprint ID"})
        return
    }
	// Delete the sprint in the database
	_, err = queries.DeleteSprint(sprintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sprint"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Sprint deleted successfully",
	})

}
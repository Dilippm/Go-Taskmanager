package controllers

import (
	"fmt"
	"net/http"

	"github.com/dilippm92/taskmanager/helpers"
	"github.com/dilippm92/taskmanager/models"
	"github.com/dilippm92/taskmanager/models/queries" // Import the queries package instead of models
	"github.com/gin-gonic/gin"
)

// TestController is a test endpoint
func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test Route"})
}

// retrieves a user by ID
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	// Call the function from the queries package to find the user
	user, err := queries.FindUserByID(userID)
	if err != nil {
		c.Error(fmt.Errorf("failed to find user with ID %s: %w", userID, err))
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)

}

// userregister
func CreateNewUser(c *gin.Context) {
	var user models.User

	// Bind the request body to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(fmt.Errorf("failed to bind request body to user model: %v", err))
		c.AbortWithStatus(http.StatusBadRequest) // Bad Request is more appropriate for validation errors
		return
	}

	// Hash the password before saving it to the database
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.Error(fmt.Errorf("failed to hash password: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError) // Internal Server Error for issues with server-side processing
		return
	}
	user.Password = hashedPassword

	// Call CreateUser to insert the user into the database
	result, err := queries.CreateUser(user)
	if err != nil {
		c.Error(fmt.Errorf("failed to create user in the database: %v", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Return a success response with the result
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"result":  result.InsertedID,
	})
}

// user login with jwttoken
func UserLogin(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the JSON input to loginData struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.Error(fmt.Errorf("failed to bind request body to user model: %v", err))
		c.AbortWithStatus(http.StatusBadRequest) // Bad Request is more appropriate for validation errors
		return
	}
	// Find the user by email
	user, err := queries.FindUserByEmail(loginData.Email)
	if err != nil {
		c.Error(fmt.Errorf("user not found"))
		c.AbortWithStatus(http.StatusInternalServerError) // Internal Server Error for issues with server-side processing
		return
	}

	// Compare the password with the stored hash using the new function
	err = helpers.ComparePasswords(user.Password, loginData.Password)
	if err != nil {
		c.Error(fmt.Errorf("email or password wrong"))
		c.AbortWithStatus(http.StatusInternalServerError) // Internal Server Error for issues with server-side processing
		return
	}
	// Generate JWT token
	token, err := helpers.GenerateJWTToken(user.ID.Hex(), user.Email)
	if err != nil {
		c.Error(fmt.Errorf("token generation failed"))
		c.AbortWithStatus(http.StatusInternalServerError) // Internal Server Error for issues with server-side processing
		return
	}
	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user})
}

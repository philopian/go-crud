package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/philopian/go-crud/initializers"
	"github.com/philopian/go-crud/model"
)

func UsersCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.Bind(&body)

	// Create a user in the DB
	user := model.User{Username: body.Username, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error})
		return
	}

	// Return the user (without password)
	c.JSON(200, gin.H{"user": user})
}

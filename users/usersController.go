package users

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	// Get data off req body
	var body struct {
		Username string
		Password string
	}
	c.Bind(&body)

	// Create a user in the DB
	user, err := createUser(body.Username, body.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func ReadAll(c *gin.Context) {
	response := getAllUsers()
	c.JSON(200, gin.H{"users": response})
}

func Read(c *gin.Context) {
	// Get the route param
	id := c.Param("id")

	// Convert id to uuid.UUID
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	// Read from database
	user, err := getUser(uuidID)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func Update(c *gin.Context) {
	// Get the route param
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	// Get data off req body
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	c.Bind(&body)

	// Read from database
	user, err := updateUser(uuidID, body.Username, body.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"user": user})

}

func Delete(c *gin.Context) {
	// Get the route param
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	deleteUser(uuidID)
	c.JSON(200, gin.H{"deleted": id})
}

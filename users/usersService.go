package users

import (
	"github.com/google/uuid"
	"github.com/philopian/go-crud/initializers"
	"github.com/philopian/go-crud/models"
	"golang.org/x/crypto/bcrypt"
)

func createUser(Username string, Password string) (*models.User, error) {
	// Bcrypt the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create a user in the DB
	uuid := uuid.New()
	user := models.User{ID: uuid, Username: Username, Password: string(hashPassword)}
	result := initializers.DB.Create(&user)

	// TODO: remove the password from the response
	return &user, result.Error
}

func getAllUsers() []models.User {
	var users []models.User
	initializers.DB.Find(&users)

	// TODO: remove the password from the response
	return users
}

func getUser(Id uuid.UUID) (*models.User, error) {
	// Get single user in the DB
	var user models.User
	result := initializers.DB.First(&user, Id)

	return &user, result.Error
}

func updateUser(Id uuid.UUID, Username string, Password string) (*models.User, error) {
	// Bcrypt the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var user models.User
	initializers.DB.First(&user, Id)

	// Update a user in the DB
	updatedUser := initializers.DB.Model(&user).Updates(models.User{Username: Username, Password: string(hashPassword)})

	return &user, updatedUser.Error
}

func deleteUser(Id uuid.UUID) error {
	// Delete a user in the DB
	result := initializers.DB.Delete(&models.User{}, Id)

	return result.Error
}

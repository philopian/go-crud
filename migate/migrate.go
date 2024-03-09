package main

import (
	"github.com/philopian/go-crud/initializers"
	"github.com/philopian/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}

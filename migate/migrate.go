package main

import (
	"github.com/philopian/go-crud/initializers"
	"github.com/philopian/go-crud/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.User{})
}

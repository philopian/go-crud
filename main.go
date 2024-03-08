package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philopian/go-crud/controllers"
	"github.com/philopian/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/users", controllers.UsersCreate)

	r.Run()
}

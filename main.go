package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philopian/go-crud/initializers"
	"github.com/philopian/go-crud/users"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/users", users.Create)
	r.GET("/users", users.ReadAll)
	r.GET("/users/:id", users.Read)
	r.PUT("/users/:id", users.Update)
	r.DELETE("/users/:id", users.Delete)

	r.Run()
}

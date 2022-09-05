package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/controllers"
	"github.com/robbyklein/pages/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
	initializers.CreatePeople()
}

func main() {
	// Create gin app
	app := gin.Default()

	// Configure app
	app.LoadHTMLGlob("templates/**/*")
	app.Static("/assets", "./assets")

	// Routing
	app.GET("/", controllers.PeopleIndexGET)

	// Start app
	app.Run()
}

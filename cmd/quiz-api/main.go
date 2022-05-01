package main

import (
	"quiz-api/controller"
	"quiz-api/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	model.ConnectDatabase()

	generalController := controller.NewController()
	// Routes
	r.GET("/quizes", generalController.GetAll(model.Quiz{}))

	// Run the server
	r.Run()
}

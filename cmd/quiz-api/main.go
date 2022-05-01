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

	quizController := controller.NewQuizController()
	// Routes
	r.GET("/quizes", quizController.GetAll)

	// Run the server
	r.Run()
}

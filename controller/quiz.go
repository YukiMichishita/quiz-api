package controller

import (
	"net/http"
	"quiz-api/model"

	"github.com/gin-gonic/gin"
)

type QuizController interface {
	GetAll(c *gin.Context)
	// PostQuiz(w http.ResponseWriter, r *http.Request)
	// PutQuiz(w http.ResponseWriter, r *http.Request)
	// DeleteQuiz(w http.ResponseWriter, r *http.Request)
}

type quizController struct {
}

func NewQuizController() controller {
	return &Controller{}
}

func (*quizController) GetAll(c *gin.Context) {
	var quizes []model.Quiz
	model.DB.Find(&quizes)
	c.JSON(http.StatusOK, gin.H{"data": quizes})
}

func (*quizController) GetById(c *gin.Context) {
	var quizes []model.Quiz
	model.DB.Find(&quizes)
	c.JSON(http.StatusOK, gin.H{"data": quizes})
}

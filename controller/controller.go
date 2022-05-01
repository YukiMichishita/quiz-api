package controller

import (
	"net/http"
	"quiz-api/model"

	"github.com/gin-gonic/gin"
)

type controller interface {
	GetAll(target_model interface{}) gin.HandlerFunc
	GetById(c *gin.Context, target_model interface{}, id uint)
	Post(c *gin.Context, target_model interface{})
	Put(c *gin.Context, target_model interface{}, id uint)
	Delete(c *gin.Context, target_model interface{}, id uint)
}

type Controller struct {
}

func NewController() controller {
	return &Controller{}
}

func getModelSlice(target_model interface{}) interface{} {
	switch target_model.(type) {
	case model.Quiz:
		return []model.Quiz{}
	case model.QuizCollectionHeader:
		return []model.QuizCollectionHeader{}
	default:
		return nil
	}
}

func (controller *Controller) GetAll(target_model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := getModelSlice(target_model)
		model.DB.Find(&data)
		c.JSON(http.StatusOK, gin.H{"data": target_model})
	}
}

func (controller *Controller) GetById(c *gin.Context, target_model interface{}, id uint) {
	model.DB.Find(target_model)
	c.JSON(http.StatusOK, gin.H{"data": target_model})
}

func (controller *Controller) Post(c *gin.Context, target_model interface{}) {
	model.DB.Find(target_model)
	c.JSON(http.StatusOK, gin.H{"data": target_model})
}

func (controller *Controller) Put(c *gin.Context, target_model interface{}, id uint) {
	model.DB.Find(target_model)
	c.JSON(http.StatusOK, gin.H{"data": target_model})
}

func (controller *Controller) Delete(c *gin.Context, target_model interface{}, id uint) {
	model.DB.Find(target_model)
	c.JSON(http.StatusOK, gin.H{"data": target_model})
}

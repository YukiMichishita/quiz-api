package controller

import (
	"net/http"
	"quiz-api/model"

	"github.com/gin-gonic/gin"
)

type controller interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context, id uint)
	Post(c *gin.Context)
	Put(c *gin.Context, id uint)
	Delete(c *gin.Context, id uint)
}

type Controller struct {
	target_model []model.Model
}

func NewController(target_model []model.Model) controller {
	return &Controller{target_model}
}

func (controller *Controller) GetAll(c *gin.Context) {
	var data []model.Model
	copy(data, controller.target_model)
	model.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (controller *Controller) GetById(c *gin.Context, id uint) {
	var data []model.Model
	copy(data, controller.target_model)
	model.DB.Find(&data, id)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (controller *Controller) Post(c *gin.Context) {
	var data []model.Model
	copy(data, controller.target_model)
	model.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (controller *Controller) Put(c *gin.Context, id uint) {
	var data []model.Model
	copy(data, controller.target_model)
	model.DB.Find(&data, id)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (controller *Controller) Delete(c *gin.Context, id uint) {
	var data []model.Model
	copy(data, controller.target_model)
	model.DB.Find(&data, id)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

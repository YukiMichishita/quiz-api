package model

import (
	"gorm.io/gorm"
)

type QuizCollectionHeader struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

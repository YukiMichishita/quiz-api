package model

import (
	"gorm.io/gorm"
)

type QuizCollectionBody struct {
	gorm.Model
	QuizCollectionHeaderId uint `json:"quiz_collection_header_id"`
	QuizCollectionHeader   QuizCollectionHeader
	QuizId                 uint `json:"quiz_id"`
	Quiz                   Quiz
	Order                  int `json:"order"`
}

package entity

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	QuizCollectionHeaderId uint `json:"quiz_collection_header_id"`
	QuizCollectionHeader   QuizCollectionHeader
	Name                   string `json:"name"`
}

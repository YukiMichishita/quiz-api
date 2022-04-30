package entity

import (
	"gorm.io/gorm"
)

type ScoreHeader struct {
	gorm.Model
	QuizCollectionHeaderId uint `json:"quiz_collection_header_id"`
	QuizCollectionHeader   QuizCollectionHeader
	UserId                 uint `json:"user_id"`
	User                   User
	Description            int `json:"description"`
}

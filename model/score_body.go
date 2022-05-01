package model

import (
	"gorm.io/gorm"
)

type ScoreBody struct {
	gorm.Model
	QuizCollectionHeaderId uint `json:"quiz_collection_header_id"`
	QuizCollectionHeader   QuizCollectionHeader
	QuizId                 uint `json:"quiz_id"`
	Quiz                   Quiz
	UserId                 uint `json:"user_id"`
	User                   User
	ScoreHeaderId          uint `json:"score_header_id"`
	ScoreHeader            ScoreHeader
	Correct                bool `json:"correct"`
}

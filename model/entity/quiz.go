package entity

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Sentence   string `json:"sentence"`
	Importance int    `json:"importance"`
	Comment    string `json:"comment"`
}

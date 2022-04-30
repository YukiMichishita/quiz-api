package main

import (
	"fmt"
	"net/http"
	"quiz-api/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	migrate()
	server := http.Server{
		Addr: ":8080",
	}
	//http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}

func migrate() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "quiz-password", "quiz-api-db:3306", "quiz",
	)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entity.User{}, &entity.Quiz{}, &entity.QuizCollectionHeader{}, &entity.QuizCollectionBody{}, &entity.Tag{}, &entity.ScoreHeader{}, &entity.ScoreBody{})
}

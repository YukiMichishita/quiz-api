package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	//mysqlの起動待ち。todo:3秒は適当なのでちゃんと起動確認するように変更する。
	time.Sleep(3 * time.Second)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "quiz-password", "quiz-api-db:3306", "quiz",
	)
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Quiz{}, &QuizCollectionHeader{}, &QuizCollectionBody{}, &Tag{}, &ScoreHeader{}, &ScoreBody{})
	DB = db
}

package main

import (
	"fmt"
	"server/internal/db"
	"server/internal/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Article{}, &model.Quiz{}, &model.QuizOption{})
}

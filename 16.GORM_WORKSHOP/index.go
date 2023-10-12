package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB = connectionDatabase()
	todoList := todov1{Username: "admin", Title: "Angular framework", Message: "Study about ngIf"}

	db.Create(&todoList)
}

func connectionDatabase() *gorm.DB {
	database, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	database.AutoMigrate(&todov1{})
	return database
}

type todov1 struct {
	gorm.Model
	Username string
	Title    string
	Message  string
}

type todov2 struct {
	ID       uint   `json:"id gorm:"primary_key`
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

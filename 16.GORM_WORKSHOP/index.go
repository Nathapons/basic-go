package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// var db *gorm.DB = connectionDatabase()
	var db *gorm.DB = connectionMySql()
	todoList := todov1{Username: "admin", Title: "Angular framework", Message: "Study about ngIf"}

	// Insert data
	fmt.Println("-------------------------- Insert --------------------------")
	db.Create(&todoList)

	// Query
	fmt.Println("-------------------------- Query --------------------------")
	query(db)

	// Update
	fmt.Println("-------------------------- Insert --------------------------")
	var todoUpdate todov1
	db.First(&todoUpdate, 1)
	update(db, todoUpdate)
	query(db)

	fmt.Println("-------------------------- Delete --------------------------")
	var deleteTmp todov1
	db.Find(&deleteTmp, "message like ?", "%For")
	delete(db, deleteTmp)
	query(db)
}

func query(_db *gorm.DB) {
	var todos []todov1
	_db.Find(&todos)
	printJson(todos)
}

func printJson(data []todov1) {
	jsonDatas, _ := json.MarshalIndent(data, " ", "")
	fmt.Println(string(jsonDatas))
}

func update(_db *gorm.DB, todo todov1) {
	_db.Model(todo).Update("Message", "Study about ngFor")
}

func delete(_db *gorm.DB, todo todov1) {
	// Soft delete
	_db.Delete(&todo)

	// Hard delete
	// _db.Unscoped().Delete(&todo)
}

func connectionDatabase() *gorm.DB {
	database, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	database.AutoMigrate(&todov1{})
	return database
}

func connectionMySql() *gorm.DB {
	dsn := "user:12345678@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
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

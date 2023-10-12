package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

func (p Product) getDetail() string {
	return fmt.Sprintf("Code %s = %d Bath", p.Code, p.Price)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key

	// productJson, _ := json.MarshalIndent(product, "", " ")
	// fmt.Println(string(productJson))

	// db.First(&product, "code = ?", "F42")
	// fmt.Println(product)

	var products []Product
	db.Find(&products, "code = ?", "D42")
	productJson2, _ := json.MarshalIndent(products, "", " ")
	fmt.Println(string(productJson2))

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 1)
}

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // many to many relation
}

type Product struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"` // many to many relation
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpertgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// cria tabela se baseando na struct
	db.AutoMigrate(&Product{}, &Category{})

	// // create cateogry
	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// category2 := Category{Name: "Eletronicos"}
	// db.Create(&category2)

	// // create product
	// product := Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	Categories: []Category{category, category2}, // many to many relation
	// }
	// db.Create(&product)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// looping
	for _, cat := range categories {
		for _, prod := range cat.Products {
			fmt.Println("Product: ", prod.Name, " - ", cat.Name)
		}
	}
}

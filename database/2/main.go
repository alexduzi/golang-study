package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
}

func main() {
	// SimpleGormOperations()

	dsn := "root:root@tcp(localhost:3306)/goexpertgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// // cria tabela se baseando na struct
	// db.AutoMigrate(&Product{}, &Category{})

	// // create cateogry
	// category := Category{Name: "Eletronics"}
	// db.Create(&category)

	// // create product
	// product := Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// }
	// db.Create(&product)

	// carregando as categorias com Preload
	var products []Product
	db.Preload("Category").Find(&products)
	for _, prod := range products {
		fmt.Printf("%+v \n", prod)
		fmt.Println()
		fmt.Printf("%+v \n", prod.Category)
	}
}

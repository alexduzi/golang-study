package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpertgorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// cria tabela se baseando na struct
	db.AutoMigrate(&Product{})

	// // inserindo
	// db.Create(&Product{
	// 	Name:  "Playstation 5",
	// 	Price: 699.00,
	// })

	// // inserindo em batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 5000.00},
	// 	{Name: "Smartphone", Price: 2000.00},
	// 	{Name: "Aoc Monitor 27pol", Price: 1000.00},
	// }
	// db.Create(&products)

	// // retornando primeiro produto
	// var product Product
	// db.First(&product)
	// fmt.Printf("%+v \n", product)

	// // retornando o produto de id 2
	// var product2 Product
	// db.First(&product2, 2)
	// fmt.Printf("%+v \n", product2)

	// // retornando primeiro produto com cláusula where
	// var smartphone Product
	// db.First(&smartphone, "name = ?", "Smartphone")
	// fmt.Printf("%+v \n", smartphone)

	// retornando todos os produtos
	var products []Product
	db.Find(&products)
	for _, prod := range products {
		fmt.Printf("%+v \n", prod)
	}
}

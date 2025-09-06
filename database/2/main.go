package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:primaryKey`
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

	// inserindo
	db.Create(&Product{
		ID:    uuid.New().String(),
		Name:  "Playstation 5",
		Price: 699.00,
	})

	// inserindo em batch
	products := []Product{
		{ID: uuid.NewString(), Name: "Notebook", Price: 5000.00},
		{ID: uuid.NewString(), Name: "Smartphone", Price: 2000.00},
		{ID: uuid.NewString(), Name: "Aoc Monitor 27pol", Price: 1000.00},
	}
	db.Create(&products)
}

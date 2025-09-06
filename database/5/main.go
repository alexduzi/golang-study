package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin()
	var c Category

	// lock pessimista
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletronics"
	tx.Debug().Save(&c)
	tx.Commit()
}

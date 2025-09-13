package main

import (
	"net/http"

	"github.com/alexduzi/golang-study/apis/configs"
	"github.com/alexduzi/golang-study/apis/internal/entity"
	"github.com/alexduzi/golang-study/apis/internal/infra/database"
	"github.com/alexduzi/golang-study/apis/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v", cfg)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}

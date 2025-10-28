package main

import (
	"context"
	"fmt"
	"test/exercise_products_http/service"
)

// https://fakestoreapi.com/products

func main() {
	client := service.NewClientApiCall()
	
	products, err := client.GetProducts(context.Background())
	if err != nil {
		panic(err)
	}

	categoriesMap, _ := client.GetCategoriesValue(products)
	
	highestCategory := client.GetHighestCategoryValue(categoriesMap)

	fmt.Println(highestCategory)
}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/exercise_products_http/model"
)

const (
	URL_PRODUCTS string = "https://fakestoreapi.com/products"
)

type ClientApiCall struct {
}

func NewClientApiCall() *ClientApiCall {
	return &ClientApiCall{}
}

func (c *ClientApiCall) GetProducts(ctx context.Context) ([]model.Product, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", URL_PRODUCTS, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	var products []model.Product
	json.Unmarshal(body, &products)
	fmt.Println(products)
	return products, nil
}

func (c *ClientApiCall) GetCategoriesValue(products []model.Product) (map[string]float64, error) {
	categoryValue := make(map[string]float64)

	for _, prod := range products {
		categoryValue[prod.Category] = prod.Price * float64(prod.Rating.Count)
	}

	return categoryValue, nil
}

func (c *ClientApiCall) GetHighestCategoryValue(categoryValue map[string]float64) model.CategoryQuanty {
	categoriesQty := c.convertToCategoryQuanty(categoryValue)

	maxValue := float64(0)
	idx := 0

	for i, cat := range categoriesQty {
		if cat.Total > maxValue {
			maxValue = cat.Total
			idx = i
		}
	}

	return categoriesQty[idx]
}

func (c *ClientApiCall) convertToCategoryQuanty(categoryValue map[string]float64) []model.CategoryQuanty {
	categoriesQty := make([]model.CategoryQuanty, 0, len(categoryValue))

	for key, val := range categoryValue {
		categoriesQty = append(categoriesQty, model.CategoryQuanty{
			Name:  key,
			Total: val,
		})
	}

	return categoriesQty
}

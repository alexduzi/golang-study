package service

import (
	"context"
	"encoding/json"
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

	var products []model.Product
	json.Unmarshal(body, &products)

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
	maxCategory := model.CategoryQuanty{
		Name:  "",
		Total: -1,
	}

	for key := range categoryValue {
		if categoryValue[key] > maxCategory.Total {
			maxCategory.Total = categoryValue[key]
			maxCategory.Name = key
		}
	}

	return maxCategory
}

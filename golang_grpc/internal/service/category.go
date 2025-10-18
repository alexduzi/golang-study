package service

import (
	"context"

	"github.com/alexduzi/golang-study/golanggrpc/internal/database"
	"github.com/alexduzi/golang-study/golanggrpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryPb := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryPb, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	categoriesPb := make([]*pb.Category, 0, len(categories))

	for _, cat := range categories {
		categoryPb := &pb.Category{
			Id:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
		}
		categoriesPb = append(categoriesPb, categoryPb)
	}

	return &pb.CategoryList{Categories: categoriesPb}, nil
}

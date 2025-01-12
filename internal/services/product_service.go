package services

import (
	"context"
	"fmt"

	db "github.com/Sinet2000/go-eshop-console/internal/data"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
)

type ProductService struct {
	repo *db.ProductRepository
}

func NewProductService(repo *db.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) ListAllProducts(ctx context.Context) ([]entities.Product, error) {
	return s.repo.ListAll(ctx)
}

func (s *ProductService) Seed(ctx context.Context, products []entities.Product) error {
	var productsToSeed []interface{}
	for _, product := range products {
		productsToSeed = append(productsToSeed, product)
	}

	_, err := s.repo.InsertProducts(ctx, productsToSeed)
	if err != nil {
		return fmt.Errorf("error inserting products: %w", err)
	}

	return nil
}
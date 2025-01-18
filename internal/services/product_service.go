package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/utils/pagination"
	"log"

	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/Sinet2000/go-eshop-console/internal/utils/file_reader"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	repo *db.ProductRepository
}

func NewProductService(repo *db.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(newProduct *entities.Product, ctx context.Context) (*entities.Product, error) {
	return s.repo.Create(newProduct, ctx)
}

func (s *ProductService) Update(updatedProduct *entities.Product, ctx context.Context) error {
	return s.repo.Update(updatedProduct, ctx)
}

func (s *ProductService) UpdateAndReturn(updatedProduct *entities.Product, ctx context.Context) (*entities.Product, error) {
	return s.repo.UpdateAndReturn(updatedProduct, ctx)
}

func (s *ProductService) ListAllProducts(ctx context.Context) ([]entities.Product, error) {
	return s.repo.ListAll(ctx)
}

func (s *ProductService) ListAllProductsPaged(ctx context.Context, pq *pagination.PageQuery) (pagination.PagedResult[entities.Product], error) {
	return s.repo.ListPaged(ctx, pq)
}

func (s *ProductService) GetById(ctx context.Context, id string) (*entities.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	product, err := s.repo.GetById(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("product not found: %v", err)
	}

	return product, nil
}

func (s *ProductService) GetProductsTotalCount(ctx context.Context) (int64, error) {
	return s.repo.CountProducts(ctx, nil)
}

func (s *ProductService) Seed(ctx context.Context) error {
	// Count the products in the collection
	count, err := s.repo.CountProducts(ctx, nil)
	if err != nil {
		log.Fatalf("Error counting products: %v", err)
	}

	logger.PrintlnColoredText("Seeding products from JSON ...", logger.GrayColor)

	if count > 0 {
		logger.PrintColoredText("Products are already seeded to DB: ", logger.GrayColor)
		return nil
	}

	fsr := &file_reader.FileSystemReader{}
	products, err := readProductsFromFile(ctx, fsr)
	if err != nil {
		return fmt.Errorf("error reading products from file: %w", err)
	}

	var productsToSeed []interface{}
	for _, product := range products {
		productsToSeed = append(productsToSeed, product)
	}

	_, err = s.repo.InsertProducts(ctx, productsToSeed)
	if err != nil {
		return fmt.Errorf("error inserting products: %w", err)
	}

	logger.PrintColoredText("The DB is seeded with the products ...", logger.GrayColor)

	return nil
}

func readProductsFromFile(ctx context.Context, reader file_reader.FileReader) ([]entities.Product, error) {
	filePaths := config.NewFilePaths()
	fileContent, err := reader.ReadFile(filePaths.ProductsFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var products []entities.Product
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return products, nil
}

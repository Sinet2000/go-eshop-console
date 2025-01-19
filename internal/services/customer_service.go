package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/Sinet2000/go-eshop-console/internal/utils/file_reader"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/internal/utils/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerService struct {
	repo *db.CustomerRepository
}

func NewCustomerService(repo *db.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) Create(newCustomer *entities.Customer, ctx context.Context) (*entities.Customer, error) {
	return s.repo.Create(newCustomer, ctx)
}

func (s *CustomerService) Update(updatedCustomer *entities.Customer, ctx context.Context) error {
	return s.repo.Update(updatedCustomer, ctx)
}

func (s *CustomerService) DeleteById(id string, ctx context.Context) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	return s.repo.DeleteById(objectID, ctx)
}

func (s *CustomerService) ListAll(ctx context.Context) ([]entities.Customer, error) {
	return s.repo.ListAll(ctx)
}

func (s *CustomerService) ListAllPaged(ctx context.Context, pq *pagination.PageQuery) (pagination.PagedResult[entities.Customer], error) {
	return s.repo.ListPaged(ctx, pq)
}

func (s *CustomerService) GetById(ctx context.Context, id string) (*entities.Customer, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	product, err := s.repo.GetById(ctx, objectID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %v", err)
	}

	return product, nil
}

func (s *CustomerService) Seed(ctx context.Context) error {
	count, err := s.repo.CountCustomers(ctx, nil)
	if err != nil {
		return fmt.Errorf("error counting customers: %w", err)
	}

	logger.PrintlnColoredText("Seeding customers from JSON ...", logger.GrayColor)

	if count > 0 {
		logger.PrintColoredText("Customers are already seeded to DB: ", logger.GrayColor)
		return nil
	}

	fsr := &file_reader.FileSystemReader{}
	customers, err := s.readCustomersFromFile(fsr, ctx)
	if err != nil {
		return fmt.Errorf("error reading customers from file: %w", err)
	}

	var customersToSeed []interface{}
	for _, product := range customers {
		customersToSeed = append(customersToSeed, product)
	}

	_, err = s.repo.InsertCustomers(ctx, customersToSeed)
	if err != nil {
		return fmt.Errorf("error inserting customers: %w", err)
	}

	logger.PrintColoredText("The DB is seeded with the customers ...", logger.GrayColor)

	return nil
}

func (s *CustomerService) readCustomersFromFile(reader file_reader.FileReader, ctx context.Context) ([]entities.Customer, error) {
	filePaths := config.NewFilePaths()
	fileContent, err := reader.ReadFile(filePaths.CustomersFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var customers []entities.Customer
	err = json.Unmarshal(fileContent, &customers)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return customers, nil
}

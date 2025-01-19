package db

import (
	"context"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/Sinet2000/go-eshop-console/internal/utils/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type CustomerRepository struct {
	collection *mongo.Collection
	*BaseRepoImpl[entities.Customer]
}

func NewCustomerRepository(db *mongo.Database) *CustomerRepository {
	return &CustomerRepository{
		collection:   db.Collection("customers"),
		BaseRepoImpl: NewBaseRepoImpl[entities.Customer]("customers", db),
	}
}

func (r *CustomerRepository) GetById(ctx context.Context, id primitive.ObjectID) (*entities.Customer, error) {
	return r.BaseRepoImpl.GetById(ctx, id)
}

func (r *CustomerRepository) Create(newCustomer *entities.Customer, ctx context.Context) (*entities.Customer, error) {
	_, err := r.collection.InsertOne(ctx, newCustomer)
	if err != nil {
		return nil, fmt.Errorf("failed to insert customer: %w", err)
	}

	return newCustomer, nil
}

func (r *CustomerRepository) Update(updatedCustomer *entities.Customer, ctx context.Context) error {
	if updatedCustomer.ID.IsZero() {
		return fmt.Errorf("customer ID is zero")
	}

	filter := bson.M{"_id": updatedCustomer.ID}
	update := bson.M{
		"$set": bson.M{
			"first_name":    updatedCustomer.FirstName,
			"last_name":     updatedCustomer.LastName,
			"company_name":  updatedCustomer.CompanyName,
			"customer_type": updatedCustomer.CustomerType,
			"contact_info":  updatedCustomer.ContactInfo,
			"address":       updatedCustomer.Address,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update customer with ID %s: %w", updatedCustomer.ID.Hex(), err)
	}

	// Check if no document was updated
	if result.MatchedCount == 0 {
		return fmt.Errorf("no customer found with ID %s", updatedCustomer.ID.Hex())
	}

	return nil
}

func (r *CustomerRepository) DeleteById(id primitive.ObjectID, ctx context.Context) error {
	filter := bson.M{"_id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete customer with ID %s: %w", id.Hex(), err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no customer found with ID %s", id.Hex())
	}

	return nil
}

func (r *CustomerRepository) ListPaged(ctx context.Context, pq *pagination.PageQuery) (pagination.PagedResult[entities.Customer], error) {
	filter, ok := pq.Filter.(bson.M)
	if !ok || filter == nil {
		filter = bson.M{}
	}

	totalCount, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Printf("Error counting documents: %v", err)
		return pagination.CreatePagedResult[entities.Customer](nil, 0, pq.PageIndex, pq.PageSize), err
	}

	skip := (pq.PageIndex - 1) * pq.PageSize
	searchOpts := options.Find().SetSkip(skip).SetLimit(pq.PageSize)
	cursor, err := r.collection.Find(ctx, filter, searchOpts)
	if err != nil {
		log.Printf("Error finding documents: %v", err)
		return pagination.CreatePagedResult[entities.Customer](nil, totalCount, pq.PageIndex, pq.PageSize), err
	}
	defer cursor.Close(ctx)

	var customers []entities.Customer
	if err := cursor.All(ctx, &customers); err != nil {
		log.Printf("Error decoding documents: %v", err)
		return pagination.CreatePagedResult[entities.Customer](nil, totalCount, pq.PageIndex, pq.PageSize), err
	}

	// Construct and return the PagedResult
	return pagination.CreatePagedResult[entities.Customer](customers, totalCount, pq.PageIndex, pq.PageSize), nil
}

func (r *CustomerRepository) ListAll(ctx context.Context) ([]entities.Customer, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Error finding customers: %v", err)
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}(cursor, ctx)

	var customers []entities.Customer
	for cursor.Next(ctx) {
		var customer entities.Customer
		if err := cursor.Decode(&customer); err != nil {
			log.Printf("Error decoding document: %v", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor iteration error: %v", err)
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) InsertCustomers(ctx context.Context, customers []interface{}) ([]interface{}, error) {
	insertResult, err := r.collection.InsertMany(ctx, customers)
	if err != nil {
		log.Printf("Error inserting customers: %v", err)
		return nil, err
	}
	return insertResult.InsertedIDs, nil
}

func (r *CustomerRepository) CountCustomers(ctx context.Context, filter interface{}) (int64, error) {
	if filter == nil {
		filter = bson.D{}
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Printf("Error counting documents in 'products' collection: %v", err)
		return 0, err
	}
	return count, nil
}

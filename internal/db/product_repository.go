package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/utils/pagination"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository provides methods to interact with the "products" collection
type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *ProductRepository) CountProducts(ctx context.Context, filter interface{}) (int64, error) {
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

func (r *ProductRepository) InsertProducts(ctx context.Context, products []interface{}) ([]interface{}, error) {
	insertResult, err := r.collection.InsertMany(ctx, products)
	if err != nil {
		log.Printf("Error inserting products: %v", err)
		return nil, err
	}
	return insertResult.InsertedIDs, nil
}

func (r *ProductRepository) ListPaged(pq *pagination.PageQuery) (pagination.PagedResult[entities.Product], error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter, ok := pq.Filter.(bson.M)
	if !ok || filter == nil {
		filter = bson.M{}
	}

	totalCount, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Printf("Error counting documents: %v", err)
		return pagination.CreatePagedResult[entities.Product](nil, 0, pq.PageIndex, pq.PageSize), err
	}

	skip := (pq.PageIndex - 1) * pq.PageSize
	searchOpts := options.Find().SetSkip(skip).SetLimit(pq.PageSize)
	cursor, err := r.collection.Find(ctx, filter, searchOpts)
	if err != nil {
		log.Printf("Error finding documents: %v", err)
		return pagination.CreatePagedResult[entities.Product](nil, totalCount, pq.PageIndex, pq.PageSize), err
	}
	defer cursor.Close(ctx)

	var products []entities.Product
	if err := cursor.All(ctx, &products); err != nil {
		log.Printf("Error decoding documents: %v", err)
		return pagination.CreatePagedResult[entities.Product](nil, totalCount, pq.PageIndex, pq.PageSize), err
	}

	// Construct and return the PagedResult
	return pagination.CreatePagedResult[entities.Product](products, totalCount, pq.PageIndex, pq.PageSize), nil
}

func (r *ProductRepository) ListAll() ([]entities.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Error finding documents: %v", err)
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}(cursor, ctx)

	var products []entities.Product
	for cursor.Next(ctx) {
		var product entities.Product
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Error decoding document: %v", err)
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor iteration error: %v", err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetById(ctx context.Context, id primitive.ObjectID) (*entities.Product, error) {
	var product entities.Product
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &product, nil
}

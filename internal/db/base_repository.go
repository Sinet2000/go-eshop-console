package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type BaseRepository[T any] interface {
	Create(entity T) error
	Update(id int, entity T) error
	GetByID(id int) (T, error)
	ListAll() ([]T, error)
	Delete(id int) error
}

type BaseRepoImpl[T any] struct {
	collection *mongo.Collection
}

func NewBaseRepoImpl[T any](collection string, db *mongo.Database) *BaseRepoImpl[T] {
	return &BaseRepoImpl[T]{
		collection: db.Collection(collection),
	}
}

func (r *BaseRepoImpl[T]) GetById(ctx context.Context, id primitive.ObjectID) (*T, error) {
	var entity T
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			entityType := fmt.Sprintf("%T", entity)
			log.Printf("Error: %s with ID %s is not found", entityType, id.Hex())
			return nil, fmt.Errorf("%s with ID %s not found", entityType, id.Hex())
		}
		return nil, err
	}

	return &entity, nil
}

func (r *BaseRepoImpl[T]) Create(entity T) error {
	return fmt.Errorf("create method not implemented for %T", entity)
}

func (r *BaseRepoImpl[T]) Update(id int, entity T) error {
	return fmt.Errorf("update method not implemented for %T", entity)
}

func (r *BaseRepoImpl[T]) ListAll() ([]T, error) {
	return nil, fmt.Errorf("ListAll method not implemented for type %T", *new(T))
}

func (r *BaseRepoImpl[T]) Delete(id int) error {
	return fmt.Errorf("delete method not implemented for type %T", *new(T))
}

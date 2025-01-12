// db/client.go
package db

import (
	"context"
	"log"

	"github.com/Sinet2000/go-eshop-console/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDb handles MongoDB connection logic
func ConnectToDb() (*mongo.Client, context.Context) {
	config.LoadConfig()

	mongoURI := config.GetEnv("MONGO_URI")
	mongoUser := config.GetEnv("MONGO_USER")
	mongoPassword := config.GetEnv("MONGO_PASSWORD")
	mongoAuthSource := config.GetEnv("MONGO_AUTH_SOURCE")

	credential := options.Credential{
		AuthSource: mongoAuthSource,
		Username:   mongoUser,
		Password:   mongoPassword,
	}

	clientOptions := options.Client().ApplyURI(mongoURI).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	return client, context.TODO()
}

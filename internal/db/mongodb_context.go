package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbContext struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoService(dbName string, ctx context.Context) (*MongoDbContext, error) {
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

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetAuth(credential).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	if err := ensureHealthy(client, ctx); err != nil {
		// If ping fails, disconnect to avoid leaving the client in an inconsistent state
		_ = client.Disconnect(context.Background())
		return nil, err
	}

	logger.PrintlnColoredText("Successfully connected to MongoDB!", logger.SuccessColor)
	return &MongoDbContext{
		Client: client,
		DB:     client.Database(dbName),
	}, nil
}

func (m *MongoDbContext) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		log.Printf("Error while disconnecting from MongoDB: %v", err)
		return err
	}

	logger.PrintlnColoredText("Disconnected from MongoDB!", logger.GrayColor)
	return nil
}

func ensureHealthy(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	logger.PrintlnColoredText("Database is healthy!", logger.SuccessColor)
	return nil
}

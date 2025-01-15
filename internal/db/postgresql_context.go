package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/jackc/pgx/v5"
)

type PostgreSqlDbContext struct {
	Conn *pgx.Conn
}

func NewPgService() (*PostgreSqlDbContext, error) {
	config.LoadConfig()

	username := config.GetEnv("POSTGRES_USER")
	password := config.GetEnv("POSTGRES_PASSWORD")
	host := config.GetEnv("POSTGRES_HOST")
	port := config.GetEnv("POSTGRES_PORT")
	dbName := config.GetEnv("POSTGRES_DB")

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Fatalf("Error: Unable to connect to the database: %v", err)
		return nil, err
	}

	var firstName string
	err = conn.QueryRow(context.Background(), "select first_name from customers where id=$1", 1).Scan(&firstName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Fatalf("No record found for ID: %d", 1)
		} else {
			log.Fatalf("QueryRow failed: %v", err)
			return nil, err
		}
	}

	fmt.Println(firstName)

	logger.PrintlnColoredText(fmt.Sprintf("Connected to PostgreSQL Db : %s!", dbName), logger.GreenTxtColorCode)

	return &PostgreSqlDbContext{Conn: conn}, nil
}

func (ctx *PostgreSqlDbContext) Disconnect() error {
	if ctx.Conn != nil {
		err := ctx.Conn.Close(context.Background())
		if err != nil {
			log.Fatalf("Error disconnecting from database: %v\n", err)
			return err
		}
		fmt.Println("Disconnected from database")
	}
	return nil
}

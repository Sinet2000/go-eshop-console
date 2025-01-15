package db

import (
	"context"
	"fmt"
	"os"
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
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	var firstName string
	err = conn.QueryRow(context.Background(), "select first_name from customers where id=$1", 1).Scan(&firstName)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Fprintf(os.Stderr, "No record found for ID: %d\n", 1)
		} else {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
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
			fmt.Fprintf(os.Stderr, "Error disconnecting from database: %v\n", err)
			return err
		}
		fmt.Println("Disconnected from database")
	}
	return nil
}

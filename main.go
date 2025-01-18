package main

import (
	"context"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/handlers"
	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"log"
)

func main() {
	isAdmin, err := utils.Confirm("Are you admin?")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if !isAdmin {
		logger.PrintlnColoredText("UNAUTHORISED", logger.ErrorColor)
		return
	}

	config.LoadConfig()
	//_, err = db.NewPgService()
	// if err != nil {
	//	log.Fatalf("Failed to connect to PostgreSQL Db: %v", err)
	// }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoDbContext, err := db.NewMongoService(config.GetEnv("MONGO_DB_NAME"), ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := mongoDbContext.Close(); err != nil {
			log.Printf("Error during MongoDB closure: %v", err)
		}
	}()

	productRepo := db.NewProductRepository(mongoDbContext.DB)

	fmt.Println()

	for {
		adminMenuHandler := handlers.NewAdminHandler(productRepo)
		isExit := adminMenuHandler.RunAdminMenu(ctx)

		if isExit {
			logger.PrintlnColoredText("Quit 🚪", logger.SuccessColor)
			fmt.Println("Goodbye! 👋")
			break
		}
	}
}

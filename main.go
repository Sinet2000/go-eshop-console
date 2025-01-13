package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/tables"
	"github.com/Sinet2000/go-eshop-console/views"
)

const productsFilePath = "data/products.json"

func main() {
	_, err := db.NewPgService()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL Db: %v", err)
	}

	mongoDbContext, err := db.NewMongoService(config.GetEnv("MONGO_DB_NAME"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := mongoDbContext.Close(); err != nil {
			log.Printf("Error during MongoDB closure: %v", err)
		}
	}()

	productRepo := db.NewProductRepository(mongoDbContext.DB)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productService := services.NewProductService(productRepo)

	adminName := "root"

	fmt.Println()
	fmt.Println()
	for {
		currentTime := time.Now().Format("2006-01-02 15:04")

		fmt.Println("WSC - Product Management 🛠️")
		fmt.Printf("Hello %s - %s\n", adminName, currentTime)
		views.DisplayMainMenu()

		var choice int
		fmt.Printf("\nSelect an option: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			logger.PrintlnColoredText("❗ Invalid input. Please enter a number between 0 and 5. ❗", logger.RedTxtColorCode)
			continue
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()

		switch choice {
		case 0:
			logger.PrintlnColoredText("🛑 Quit", logger.GreenTxtColorCode)
			fmt.Println("Goodbye! 👋")

			return
		case 1:
			logger.PrintlnColoredText("📜 List Products", logger.GreenTxtColorCode)

			productStock, err := productService.ListAllProducts(ctx)
			if err != nil {
				log.Fatalf("Error fetching products: %v", err)
			}

			tables.ListProducts(productStock)
		case 2:

			var productID string
			fmt.Printf("\nEnter the product ID:")
			_, err = fmt.Scan(&productID)
			if err != nil {
				logger.PrintlnColoredText("❗ Invalid input. Please enter valid product ID ❗", logger.RedTxtColorCode)
				continue
			}

			productDetails, err := productService.GetProductById(ctx, productID)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			views.DisplayProductDetails(productDetails)
		case 6:
			productService.Seed(ctx)
		default:
			fmt.Println("❗Invalid choice. Please try again. ❗")
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}
}

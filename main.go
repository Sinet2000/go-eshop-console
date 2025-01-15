package main

import (
	"fmt"
	"log"

	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/handlers"
	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
)

func main() {
	isAdmin, err := utils.Confirm("Are you admin?")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if !isAdmin {
		logger.PrintlnColoredText("UNAUTHORISED", logger.RedTxtColorCode)
		return
	}

	config.LoadConfig()
	//_, err = db.NewPgService()
	//if err != nil {
	//	log.Fatalf("Failed to connect to PostgreSQL Db: %v", err)
	//}

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
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// productService := services.NewProductService(productRepo)

	fmt.Println()

	for {
		adminMenuHandler := handlers.NewAdminHandler(productRepo)
		shouldExit := adminMenuHandler.RunAdminMenu()

		// 	return
		// case 1:
		// 	logger.PrintlnColoredText("📜 List Products", logger.GreenTxtColorCode)

		// 	productStock, err := productService.ListAllProducts(ctx)
		// 	if err != nil {
		// 		log.Fatalf("Error fetching products: %v", err)
		// 	}

		// 	tables.ListProducts(productStock)
		// case 2:

		// 	var productID string
		// 	fmt.Printf("\nEnter the product ID:")
		// 	_, err = fmt.Scan(&productID)
		// 	if err != nil {
		// 		logger.PrintlnColoredText("❗ Invalid input. Please enter valid product ID ❗", logger.RedTxtColorCode)
		// 		continue
		// 	}

		// 	productDetails, err := productService.GetProductById(ctx, productID)
		// 	if err != nil {
		// 		fmt.Println("Error:", err)
		// 		continue
		// 	}

		// 	views.DisplayProductDetails(productDetails)
		// case 6:
		// 	productService.Seed(ctx)
		// default:
		// 	fmt.Println("❗Invalid choice. Please try again. ❗")
		// }

		if shouldExit {
			logger.PrintlnColoredText("Quit 🚪", logger.GreenTxtColorCode)
			fmt.Println("Goodbye! 👋")
			break
		}
	}
}

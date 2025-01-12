package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/tables"
	"github.com/Sinet2000/go-eshop-console/views"
)

func main() {
	dbClient, ctx := db.ConnectToDb()
	defer func() {
		if err := dbClient.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	productRepo := db.NewProductRepository(dbClient.Database(os.Getenv("MONGO_DB_NAME")))
	productService := services.NewProductService(productRepo)

	// Fetch and display all products
	productStock, err := productService.ListAllProducts(ctx)
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}

	adminName := "root"

	fmt.Println()
	fmt.Println()
	for {
		currentTime := time.Now().Format("2006-01-02 15:04")

		fmt.Println("WSC - Product Management 🛠️")
		fmt.Printf("Hello %s - %s\n", adminName, currentTime)
		views.ShowMainMenu()

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
		case 1:
			logger.PrintlnColoredText("📜 List Products", logger.GreenTxtColorCode)
			tables.ListProducts(productStock)
		case 2:
			ShowProductDetailsById(productService, ctx)
		case 0:
			logger.PrintlnColoredText("🛑 Quit", logger.GreenTxtColorCode)
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("❗Invalid choice. Please try again. ❗")
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}
}

func ShowProductDetailsById(service *services.ProductService, ctx context.Context) {
	var inputId string
	fmt.Println("Enter the product ID to get details:")
	fmt.Scanln(&inputId)

	product, err := service.GetProductById(ctx, inputId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Show product details
	fmt.Printf("Product Details:\nID: %s\nName: %s\nSKU: %s\nPrice: %.2f\n",
		product.ID.Hex(), product.Name, product.SKU, product.Price)
}

// func addNewProduct(productStock *[]entities.Product) {
// 	newProduct, err := entities.CreateProduct(
// 		len(*productStock)-1,
// 		"Apple MacBook Pro 14-inch",
// 		"AMP14-001",
// 		"A high-performance laptop with Apple's M1 Pro chip, featuring a stunning Retina display and long-lasting battery life.",
// 		1299.99, 45, "")
// 	if err != nil {
// 		logger.PrintColoredText("❗An error occurred: ", logger.RedTxtColorCode)
// 		fmt.Println(err)
// 		return
// 	}

// 	// Append the new product to the stock
// 	*productStock = append(*productStock, *newProduct)
// }

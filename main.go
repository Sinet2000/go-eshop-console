package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	db "github.com/Sinet2000/go-eshop-console/internal/data"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/views"
)

const productsFilePath = "data/products.json"

func main() {
	dbClient, ctx := db.ConnectToDb()
	defer func() {
		if err := dbClient.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	productRepo := db.NewProductRepository(dbClient.Database(os.Getenv("MONGO_DB_NAME")))
	productService := services.NewProductService(productRepo)

	// Count the products in the collection
	count, err := productRepo.CountProducts(ctx)
	if err != nil {
		log.Fatalf("Error counting products: %v", err)
	}

	if count == 0 {
		fmt.Println("Seeding the database with products...")
		products, err := readProductsFromFile(productsFilePath)
		if err != nil {
			logger.PrintColoredText("❗An error occurred: ", logger.RedTxtColorCode)
			fmt.Println(err)
			return
		}
		err = productService.Seed(ctx, products)
		if err != nil {
			logger.PrintColoredText("❗An error occurred while seeding: ", logger.RedTxtColorCode)
			fmt.Println(err)
			return
		}
	}

	// Fetch and display all products
	productStock, err := productService.ListAllProducts(ctx)
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}

	adminName := "root"

	newProduct, err := entities.CreateProduct(
		len(productStock)-1,
		"Apple MacBook Pro 14-inch",
		"AMP14-001",
		"A high-performance laptop with Apple's M1 Pro chip, featuring a stunning Retina display and long-lasting battery life.",
		1299.99, 45, "")
	if err != nil {
		logger.PrintColoredText("❗An error occurred: ", logger.RedTxtColorCode)
		fmt.Println(err)
	}

	productStock = append(productStock, *newProduct)

	views.ShowProductTable(productStock)

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

		switch choice {
		case 1:
			logger.PrintlnColoredText("📜 List Products", logger.GreenTxtColorCode)
		case 0:
			logger.PrintlnColoredText("🛑 Quit", logger.GreenTxtColorCode)
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("❗Invalid choice. Please try again. ❗")
		}
	}
}

func readProductsFromFile(filePath string) ([]entities.Product, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the entire file content into a byte slice
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Create a slice to hold products
	var products []entities.Product

	// Unmarshal JSON into the products slice
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Return the populated slice of products
	return products, nil
}
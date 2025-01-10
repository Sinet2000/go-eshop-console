package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Sinet2000/go-eshop-console/menu"
	product_scope "github.com/Sinet2000/go-eshop-console/modules/product"
	"github.com/Sinet2000/go-eshop-console/utils/logger"
)

const productsFilePath = "data/products.json"

func main() {
	adminName := "root"
	productStock, err := readProductsFromFile(productsFilePath)
	if err != nil {
		logger.PrintColoredText("❗An error occurred: ", logger.RedTxtColorCode)
		fmt.Println(err)
	} else {
		logger.PrintlnColoredText("✅ Seeded products from json: ", logger.GreenTxtColorCode)
	}

	newProduct, err := product_scope.CreateProduct(
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

	product_scope.PrintProductTable(productStock)

	fmt.Println()
	fmt.Println()
	for {
		currentTime := time.Now().Format("2006-01-02 15:04")

		fmt.Println("WSC - Product Management 🛠️")
		fmt.Printf("Hello %s - %s\n", adminName, currentTime)
		menu.ShowMainMenu()

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

func readProductsFromFile(filePath string) ([]product_scope.Product, error) {
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
	var products []product_scope.Product

	// Unmarshal JSON into the products slice
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// Return the populated slice of products
	return products, nil
}

// Quit the program: 🛑
// Goodbye: 👋
// Delete action: 🗑️
// Create new item: ➕
// List items: 📜
// Update in progress: 🔄
// Success: ✅
// Error: ❗ (Exclamation Mark) or ⚠️ (Warning Sign)
// Management: 🛠️
// Products: 📦
// Order: 📝
// Person: 👤

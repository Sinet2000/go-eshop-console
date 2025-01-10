package main

import (
	"fmt"
	"time"

	"github.com/Sinet2000/go-eshop-console/domain/models"
	"github.com/Sinet2000/go-eshop-console/utils"
)

func main() {
	adminName := "root"
	productStock := []models.Product{}
	newProduct, err := models.CreateProduct(
		len(productStock)-1,
		"Apple MacBook Pro 14-inch",
		"AMP14-001",
		"A high-performance laptop with Apple's M1 Pro chip, featuring a stunning Retina display and long-lasting battery life.",
		1299.99, 45, "")
	if err != nil {
		utils.PrintColoredText("❗An error occurred: ", utils.RedTxtColorCode)
		fmt.Println(err)
	}

	productStock = append(productStock, *newProduct)

	for {
		fmt.Printf("Hello %s - %s\n", adminName, time.Now().Format("2006-01-02 15:04"))
		fmt.Println("WSC - Product Management 🛠️")
		fmt.Println("------------------------------------------")
		fmt.Println("[1] 📜 List Products")
		fmt.Println("[2] 📝 Product details")
		fmt.Println("[3] 🔄 Edit product")
		fmt.Println("[4] 🗑️ Delete product")
		fmt.Println("[5] 🆕 Create product")
		fmt.Println("[0] 🛑 Quit")
		fmt.Printf("\n\nSelect an option: ")

		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil {
			utils.PrintlnColoredText("❗ Invalid input. Please enter a number between 0 and 5. ❗", utils.RedTxtColorCode)
			continue
		}

		switch choice {
		case 1:
			utils.PrintlnColoredText("📜 List Products", utils.GreenTxtColorCode)
		case 0:
			utils.PrintlnColoredText("🛑 Quit", utils.GreenTxtColorCode)
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("❗Invalid choice. Please try again. ❗")
		}
	}
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

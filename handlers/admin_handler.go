package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/views"
)

type AdminHandler struct {
	productService *services.ProductService
}

func NewAdminHandler(productRepo *db.ProductRepository) *AdminHandler {
	productService := services.NewProductService(productRepo)
	return &AdminHandler{productService: productService}
}

func (h *AdminHandler) RunAdminMenu(ctx context.Context) bool {
	productHandler := NewProductMngmtHandler(h.productService)

	for {
		views.DisplayAdminMenu()

		option, err := utils.PromptIntInput("Select an option: ")
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		switch option {
		case 1:
			productHandler.HandleAdminManageProducts(ctx)
		case 2:
			handleManageOrders()
		case 3:
			handleManageCustomers()
		case 4:
			handleAnalytics()
		case 5:
			handleSystemSettings()
		case 0:
			return true
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}

}

// handleManageOrders handles actions related to order management.
func handleManageOrders() {
	fmt.Println("Opening order management... 🛒")
	// Implement logic for managing orders.
	fmt.Println("Order management not yet implemented.")
}

// handleManageCustomers handles actions related to customer management.
func handleManageCustomers() {
	fmt.Println("Opening customer management... 👥")
	// Implement logic for managing customers.
	fmt.Println("Customer management not yet implemented.")
}

// handleAnalytics handles actions related to analytics.
func handleAnalytics() {
	fmt.Println("Opening analytics... 📊")
	// Implement logic for analytics (e.g., report generation).
	fmt.Println("Analytics not yet implemented.")
}

// handleSystemSettings handles actions related to system settings.
func handleSystemSettings() {
	fmt.Println("Opening system settings... ⚙️")
	// Implement logic for updating system settings.
	fmt.Println("System settings not yet implemented.")
}

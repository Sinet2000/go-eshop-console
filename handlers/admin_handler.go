package handlers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/views"
)

type AdminHandler struct {
	productService  *services.ProductService
	customerService *services.CustomerService
}

func NewAdminHandler(dbContext *mongo.Database) *AdminHandler {
	productRepo := db.NewProductRepository(dbContext)
	productService := services.NewProductService(productRepo)

	customerRepo := db.NewCustomerRepository(dbContext)
	customerService := services.NewCustomerService(customerRepo)
	return &AdminHandler{productService: productService, customerService: customerService}
}

func (h *AdminHandler) RunAdminMenu(ctx context.Context) bool {

	for {
		views.DisplayAdminMenu()

		option, err := utils.PromptIntInput("Select an option: ")
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		switch option {
		case 1:
			productHandler := NewAdminProductMngmtHandler(h.productService)
			productHandler.HandleAdminManageProducts(ctx)
		case 2:
			handleManageOrders()
		case 3:
			fmt.Println("Opening customer management... 👥")
			customerHandler := NewAdminCustomerMgmtHandler(h.customerService)
			customerHandler.handleUpdateCustomer(ctx)
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

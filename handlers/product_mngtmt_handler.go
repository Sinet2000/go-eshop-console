package handlers

import (
	"fmt"

	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/views"
)

type ProductMngmtHandler struct {
	productService *services.ProductService
}

func NewProductMngmtHandler(productService *services.ProductService) *ProductMngmtHandler {
	return &ProductMngmtHandler{productService: productService}
}

func (h *ProductMngmtHandler) HandleAdminManageProducts() {
	views.DisplayAdminProductMngmtMenu()

	for {
		option, err := utils.PromptUserForSelection()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		switch option {
		case 0:
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

package handlers

import (
	"context"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/tables"
	"github.com/Sinet2000/go-eshop-console/views"
	"log"
)

type SharedProductHandler struct {
	productService *services.ProductService
}

func NewSharedProductHandler(productService *services.ProductService) *SharedProductHandler {
	return &SharedProductHandler{productService: productService}
}

func (h *SharedProductHandler) HandleListProducts(ctx context.Context) {
	logger.PrintlnColoredText("📜 Retrieving products from storage ...", logger.GrayColor)

	productsList, err := h.productService.ListAll(ctx)
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}

	tables.ListProducts(productsList)

	fmt.Println("\nPress Enter to continue...")
	_, err = fmt.Scanln()
	if err != nil {
		return
	}
}

func (h *SharedProductHandler) HandleGetProductDetails(ctx context.Context) *entities.Product {
	productId, err := utils.PromptStrInput("Enter the product ID: ")
	if err != nil {
		logger.PrintlnColoredText("❗ Invalid input. Please enter valid product ID ❗", logger.ErrorColor)
		return nil
	}

	productDetails, err := h.productService.GetById(ctx, productId)
	if err != nil {
		logger.PrintlnColoredText(fmt.Sprintf("❗Error. Cannot get product details [ID=%s]. Error: %v", productId, err), logger.ErrorColor)
		return nil
	}

	views.DisplayProductDetails(productDetails)

	fmt.Println("\nPress Enter to continue...")
	_, err = fmt.Scanln()
	if err != nil {
		return nil
	}

	return productDetails
}

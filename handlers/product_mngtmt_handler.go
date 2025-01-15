package handlers

import (
	"context"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/utils/pagination"
	"log"
	"os"
	"time"

	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/tables"
	"github.com/Sinet2000/go-eshop-console/views"
	"github.com/olekukonko/tablewriter"
)

type ProductMngmtHandler struct {
	productService *services.ProductService
}

func NewProductMngmtHandler(productService *services.ProductService) *ProductMngmtHandler {
	return &ProductMngmtHandler{productService: productService}
}

func (h *ProductMngmtHandler) HandleAdminManageProducts() {

	for {
		views.DisplayAdminProductMngmtMenu()
		option, err := utils.PromptIntInput()
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		switch option {
		case 1:
			h.handleListProducts()
		case 2:
			h.handleListProductsPaged()
		case 3:
			h.handleCreateProduct()
		case 4:
			h.handleUpdateProduct()
		case 5:
			h.handleDeleteProduct()
		case 9:
			h.handleSeedProducts()
		case 0:
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func (h *ProductMngmtHandler) handleListProducts() {
	logger.PrintlnColoredText("📜 Retrieving products from storage ...", logger.GrayTxtColorCode)

	productsList, err := h.productService.ListAllProducts()
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

func (h *ProductMngmtHandler) handleListProductsPaged() {
	logger.PrintlnColoredText("📜 Retrieving products from storage ...", logger.GrayTxtColorCode)

	for {
		productsPageResult, err := h.productService.ListAllProductsPaged(&pagination.PageQuery{PageIndex: 1, PageSize: 5})
		if err != nil {
			log.Fatalf("%v", err)
			return
		}

		fmt.Printf("Admin: Products (Page %d) \n", productsPageResult.Page)
		tables.ListProducts(productsPageResult.Data)

		fmt.Println("-----------------------------------")
		fmt.Printf("Total count: %d \n", productsPageResult.TotalCount)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", ""})

		table.Append([]string{"[N] Next Page ➡️", "[B] Back ⬅️"})
		table.Append([]string{"[1]", "Update product"})
		table.Append([]string{"[2]", "Product Details"})
		table.Append([]string{"[3]", "Delete Product"})
		table.Append([]string{"[0]", "Quit 🚪"})
		table.Render()

		option, err := utils.PromptStrInput()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Println("\nPress Enter to continue...")
		_, err = fmt.Scanln()
		if err != nil {
			return
		}

		switch option {
		case "N":
			fmt.Println("Going forward ➡️ ... ")
		case "B":
			fmt.Println("⬅️ Going back ... ")
		case "0":
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func (h *ProductMngmtHandler) handleCreateProduct() {

}

func (h *ProductMngmtHandler) handleUpdateProduct() {

}

func (h *ProductMngmtHandler) handleDeleteProduct() {

}

func (h *ProductMngmtHandler) handleSeedProducts() {
	logger.PrintlnColoredText("Seeding products from JSON ...", logger.GrayTxtColorCode)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	h.productService.Seed(ctx)
}

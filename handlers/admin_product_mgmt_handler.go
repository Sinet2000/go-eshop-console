package handlers

import (
	"context"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/internal/utils/pagination"
	"github.com/Sinet2000/go-eshop-console/tables"
	"github.com/Sinet2000/go-eshop-console/views"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
)

type AdminProductMngmtHandler struct {
	productService       *services.ProductService
	sharedProductHandler *SharedProductHandler
}

func NewAdminProductMngmtHandler(productService *services.ProductService) *AdminProductMngmtHandler {
	sharedProductHandler := NewSharedProductHandler(productService)
	return &AdminProductMngmtHandler{productService: productService, sharedProductHandler: sharedProductHandler}
}

func (h *AdminProductMngmtHandler) HandleAdminManageProducts(ctx context.Context) {
	for {
		views.DisplayAdminProductMngmtMenu()
		option, err := utils.PromptIntInput("Select an option: ")
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		switch option {
		case 1:
			h.sharedProductHandler.HandleListProducts(ctx)
		case 2:
			h.handleListProductsPaged(ctx)
		case 3:
			h.handleCreateProduct(ctx)
		case 4:
			h.handleUpdateProduct(ctx)
		case 5:
			h.handleDeleteProduct(ctx)
		case 6:
			h.sharedProductHandler.HandleGetProductDetails(ctx)
		case 9:
			h.handleSeedProducts(ctx)
		case 0:
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func (h *AdminProductMngmtHandler) handleListProductsPaged(ctx context.Context) {
	logger.PrintlnColoredText("📜 Retrieving products from storage ...", logger.GrayColor)
	var pageIndex int64 = 1

	for {
		productsPageResult, err := h.productService.ListAllPaged(ctx, &pagination.PageQuery{PageIndex: pageIndex, PageSize: 5})
		if err != nil {
			log.Fatalf("%v", err)
			return
		}

		fmt.Printf("Admin: Products (Page %d) \n", productsPageResult.Page)
		tables.ListProducts(productsPageResult.Data)

		fmt.Println("\n-----------------------------------")
		fmt.Printf("Total count: %d \n", productsPageResult.TotalCount)
		fmt.Println("")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", ""})

		table.Append([]string{"[N] Next Page ➡️", "[B] Back ⬅️"})
		table.Append([]string{"[1]", "Update product"})
		table.Append([]string{"[2]", "Product Details"})
		table.Append([]string{"[3]", "Delete Product"})
		table.Append([]string{"[0]", "Quit 🚪"})
		table.Render()

		option, err := utils.PromptStrInput("Select an option: ")
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		switch option {

		case "N":
			if !productsPageResult.HasNextPage {
				logger.PrintlnColoredText("Cannot go to the next page!", logger.WarningColor)
			} else {
				pageIndex++
			}

		case "B":
			if !productsPageResult.HasPrevPage {
				logger.PrintlnColoredText("Cannot go to the previous page!", logger.WarningColor)
			} else {
				pageIndex--
			}

		case "1":
			h.handleUpdateProduct(ctx)
		case "2":
			h.sharedProductHandler.HandleGetProductDetails(ctx)
		case "3":
			h.handleDeleteProduct(ctx)
		case "0":
			logger.PrintlnColoredText("Exiting Product Management ...", logger.GrayColor)
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}

		fmt.Println("\nPress Enter to continue...")
		_, err = fmt.Scanln()
		if err != nil {
			return
		}
	}
}

func (h *AdminProductMngmtHandler) handleCreateProduct(ctx context.Context) {
	fmt.Println("\n----------------------------------------------")
	fmt.Println("|               Create a Product              |")
	fmt.Println("----------------------------------------------")

	name, err := utils.PromptStrInput("Name: ")
	if err != nil {
		logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
	}

	description, err := utils.PromptStrInput("Description: ")
	if err != nil {
		logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
	}

	sku, err := utils.PromptStrInput("SKU: ")
	if err != nil {
		logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
	}

	priceStr, err := utils.PromptStrInput("Price: ")
	if err != nil {
		logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
	}

	stockStr, err := utils.PromptStrInput("Stock: ")
	if err != nil {
		logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
	}

	newProduct, err := entities.CreateProduct(name, sku, description, price, stock)
	if err != nil {
		logger.PrintlnColoredText(fmt.Sprintf("Cannot create the product object: %v", err), logger.ErrorColor)
	}

	_, err = h.productService.Create(newProduct, ctx)
	if err != nil {
		logger.PrintlnColoredText(fmt.Sprintf("%v", err), logger.ErrorColor)
	}
}

func (h *AdminProductMngmtHandler) handleDeleteProduct(ctx context.Context) {
	productId, err := utils.PromptStrInput("Enter the product ID: ")
	if err != nil {
		logger.PrintlnColoredText("❗ Invalid input. Please enter valid product ID ❗", logger.ErrorColor)
		return
	}

	err = h.productService.DeleteById(productId, ctx)
	if err != nil {
		fmt.Printf("\n%v", err)
	}
}

func (h *AdminProductMngmtHandler) handleSeedProducts(ctx context.Context) {
	err := h.productService.Seed(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func (h *AdminProductMngmtHandler) handleUpdateProduct(ctx context.Context) {
	fmt.Println("\n----------------------------------------------")
	fmt.Println("|         Update an Existing Product         |")
	fmt.Println("----------------------------------------------")

	// Print product details to update
	var productToUpdate = h.sharedProductHandler.HandleGetProductDetails(ctx)

	fmt.Println("\nWhat would you like to update?")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", ""})
	table.Append([]string{"[1]", "Name"})
	table.Append([]string{"[2]", "Description"})
	table.Append([]string{"[3]", "SKU"})
	table.Append([]string{"[4]", "Price"})
	table.Append([]string{"[5]", "Stock"})
	table.Append([]string{"[9]", "Save Changes"})
	table.Append([]string{"[0]", "Cancel"})
	table.Render()

	for {
		updateOption, err := utils.PromptIntInput("Select an option (0-9): ")
		if err != nil {
			logger.PrintlnColoredText("❗Invalid input. Please enter valid options!", logger.ErrorColor)
			return
		}

		switch updateOption {
		case 1:
			fmt.Printf("Name: %s\n", productToUpdate.Name)
			newName, err := utils.PromptStrInput("New Name: ")
			if err != nil {
				logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
			}
			productToUpdate.Name = newName
		case 2:
			fmt.Printf("Description: %s\n", productToUpdate.Description)
			newDescription, err := utils.PromptStrInput("New Description: ")
			if err != nil {
				logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
			}
			productToUpdate.Description = newDescription
		case 3:
			fmt.Printf("SKU: %s\n", productToUpdate.SKU)
			newSku, err := utils.PromptStrInput("New SKU: ")
			if err != nil {
				logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
			}
			productToUpdate.SKU = newSku
		case 4:
			fmt.Printf("Price: %.2f\n", productToUpdate.Price)
			newPriceStr, err := utils.PromptStrInput("New Price: ")
			if err == nil {
				newPrice, err := strconv.ParseFloat(newPriceStr, 64)
				if err != nil {
					logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
				} else {
					productToUpdate.Price = newPrice
				}
			}
		case 5:
			fmt.Printf("Stock: %d\n", productToUpdate.Stock)
			newStockStr, err := utils.PromptStrInput("New Stock: ")
			if err == nil {
				newStock, err := strconv.Atoi(newStockStr)
				if err != nil {
					logger.PrintlnColoredText("❗Invalid input.", logger.ErrorColor)
				} else {
					productToUpdate.Stock = newStock
				}
			}
		case 9:
			err := h.productService.Update(productToUpdate, ctx)
			if err != nil {
				fmt.Printf("Failed to update product: %v\n", err)
			} else {
				logger.PrintlnColoredText("\nSuccessfully updated product", logger.SuccessColor)
			}

			return
		case 0:
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

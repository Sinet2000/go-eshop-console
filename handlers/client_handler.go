package handlers

import (
	"context"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/services"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"github.com/Sinet2000/go-eshop-console/internal/utils/logger"
	"github.com/Sinet2000/go-eshop-console/internal/utils/pagination"
	"github.com/Sinet2000/go-eshop-console/tables"
	"github.com/Sinet2000/go-eshop-console/views"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

type ClientHandler struct {
	productService       *services.ProductService
	sharedProductHandler *SharedProductHandler
}

func NewClientHandler(productRepo *db.ProductRepository) *ClientHandler {
	productService := services.NewProductService(productRepo)
	sharedProductHandler := NewSharedProductHandler(productService)
	return &ClientHandler{productService: productService, sharedProductHandler: sharedProductHandler}
}

func (h *ClientHandler) RunClientMenu(ctx context.Context) bool {

	for {
		views.DisplayClientMenu()

		option, err := utils.PromptIntInput("Select an option: ")
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		switch option {
		case 1:
			h.handleClientShopping(ctx)
		case 0:
			return true
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func (h *ClientHandler) handleClientShopping(ctx context.Context) {
	for {
		views.DisplayClientShoppingMenu()
		option, err := utils.PromptIntInput("Select an option: ")
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		switch option {
		case 1:
			h.sharedProductHandler.HandleListProducts(ctx)
		case 2:
			h.handleListProductsPaged(ctx)
		case 3:
			h.sharedProductHandler.HandleGetProductDetails(ctx)
		case 0:
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func (h *ClientHandler) handleListProductsPaged(ctx context.Context) {
	logger.PrintlnColoredText("📜 Retrieving products from storage ...", logger.GrayColor)
	var pageIndex int64 = 1

	for {
		productsPageResult, err := h.productService.ListAllPaged(ctx, &pagination.PageQuery{PageIndex: pageIndex, PageSize: 5})
		if err != nil {
			log.Fatalf("%v", err)
			return
		}

		fmt.Printf("Products (Page %d) \n", productsPageResult.Page)
		tables.ListProducts(productsPageResult.Data)

		fmt.Println("\n-----------------------------------")
		fmt.Printf("Total count: %d \n", productsPageResult.TotalCount)
		fmt.Println("")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", ""})

		table.Append([]string{"[N] Next Page ➡️", "[B] Back ⬅️"})
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

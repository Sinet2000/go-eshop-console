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
)

type AdminCustomerMgmtHandler struct {
	customerService *services.CustomerService
}

func NewAdminCustomerMgmtHandler(customerService *services.CustomerService) *AdminCustomerMgmtHandler {
	return &AdminCustomerMgmtHandler{customerService: customerService}
}

func (h *AdminCustomerMgmtHandler) HandleAdminManageCustomers(ctx context.Context) {
	for {
		views.DisplayAdminCustomerMgmtMenu()
		option, err := utils.PromptIntInput("Select an option: ")
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}

		switch option {
		case 1:
			h.handleListCustomers(ctx)
		case 2:
			h.handleListCustomersPaged(ctx)
		case 3:
			h.handleCreateCustomer(ctx)
		case 4:
			h.handleUpdateCustomer(ctx)
		case 5:
			h.handleDeleteCustomer(ctx)
		case 6:
			h.handleGetCustomerDetails(ctx)
		case 9:
			h.handleSeedCustomers(ctx)
		case 0:
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func (h *AdminCustomerMgmtHandler) handleListCustomers(ctx context.Context) {
	logger.PrintlnColoredText("📜 Retrieving customers from storage ...", logger.GrayColor)

	productsList, err := h.customerService.ListAll(ctx)
	if err != nil {
		log.Fatalf("Error fetching customers: %v", err)
	}

	tables.ListCustomers(productsList)

	fmt.Println("\nPress Enter to continue...")
	_, err = fmt.Scanln()
	if err != nil {
		return
	}
}

func (h *AdminCustomerMgmtHandler) handleListCustomersPaged(ctx context.Context) {
	logger.PrintlnColoredText("📜 Retrieving customer from storage ...", logger.GrayColor)
	var pageIndex int64 = 1

	for {
		customersPageResult, err := h.customerService.ListAllPaged(ctx, &pagination.PageQuery{PageIndex: pageIndex, PageSize: 5})
		if err != nil {
			log.Fatalf("%v", err)
			return
		}

		fmt.Printf("Admin: Customers (Page %d) \n", customersPageResult.Page)
		tables.ListCustomers(customersPageResult.Data)

		fmt.Println("\n-----------------------------------")
		fmt.Printf("Total count: %d \n", customersPageResult.TotalCount)
		fmt.Println("")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", ""})

		table.Append([]string{"[N] Next Page ➡️", "[B] Back ⬅️"})
		table.Append([]string{"[1]", "Update customer"})
		table.Append([]string{"[2]", "Customer Details"})
		table.Append([]string{"[3]", "Delete Customer"})
		table.Append([]string{"[0]", "Quit 🚪"})
		table.Render()

		option, err := utils.PromptStrInput("Select an option: ")
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		switch option {

		case "N":
			if !customersPageResult.HasNextPage {
				logger.PrintlnColoredText("Cannot go to the next page!", logger.WarningColor)
			} else {
				pageIndex++
			}

		case "B":
			if !customersPageResult.HasPrevPage {
				logger.PrintlnColoredText("Cannot go to the previous page!", logger.WarningColor)
			} else {
				pageIndex--
			}

		case "1":
			h.handleUpdateCustomer(ctx)
		case "2":
			h.handleGetCustomerDetails(ctx)
		case "3":
			h.handleDeleteCustomer(ctx)
		case "0":
			logger.PrintlnColoredText("Exiting Customer Management ...", logger.GrayColor)
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

func (h *AdminCustomerMgmtHandler) handleCreateCustomer(ctx context.Context) {
	fmt.Println("\n----------------------------------------------")
	fmt.Println("|               Create a Customer             |")
	fmt.Println("----------------------------------------------")

	// Prompt for first name
	firstName, err := utils.PromptStrInput("First Name: ")
	if err != nil || firstName == "" {
		logger.PrintlnColoredText("❗Invalid input for First Name.", logger.ErrorColor)
		return
	}

	// Prompt for last name
	lastName, err := utils.PromptStrInput("Last Name: ")
	if err != nil || lastName == "" {
		logger.PrintlnColoredText("❗Invalid input for Last Name.", logger.ErrorColor)
		return
	}

	// Prompt for customer type
	customerTypeStr, err := utils.PromptStrInput("Customer Type (Individual/Company): ")
	if err != nil || (customerTypeStr != "Individual" && customerTypeStr != "Company") {
		logger.PrintlnColoredText("❗Invalid input for Customer Type. Must be 'Individual' or 'Company'.", logger.ErrorColor)
		return
	}

	var customerType entities.CustomerType
	if customerTypeStr == "Individual" {
		customerType = entities.Individual
	} else {
		customerType = entities.Company
	}

	// Prompt for company name if customer is a company
	var companyName string
	if customerType == entities.Company {
		companyName, err = utils.PromptStrInput("Company Name: ")
		if err != nil || companyName == "" {
			logger.PrintlnColoredText("❗Invalid input for Company Name.", logger.ErrorColor)
			return
		}
	}

	// Prompt for email
	email, err := utils.PromptStrInput("Email: ")
	if err != nil || email == "" {
		logger.PrintlnColoredText("❗Invalid input for Email.", logger.ErrorColor)
		return
	}

	// Prompt for phone number
	phone, err := utils.PromptStrInput("Phone: ")
	if err != nil || phone == "" {
		logger.PrintlnColoredText("❗Invalid input for Phone.", logger.ErrorColor)
		return
	}

	// Prompt for address details
	street, err := utils.PromptStrInput("Street: ")
	if err != nil || street == "" {
		logger.PrintlnColoredText("❗Invalid input for Street.", logger.ErrorColor)
		return
	}

	city, err := utils.PromptStrInput("City: ")
	if err != nil || city == "" {
		logger.PrintlnColoredText("❗Invalid input for City.", logger.ErrorColor)
		return
	}

	state, err := utils.PromptStrInput("State: ")
	if err != nil || state == "" {
		logger.PrintlnColoredText("❗Invalid input for State.", logger.ErrorColor)
		return
	}

	postalCode, err := utils.PromptStrInput("Postal Code: ")
	if err != nil || postalCode == "" {
		logger.PrintlnColoredText("❗Invalid input for Postal Code.", logger.ErrorColor)
		return
	}

	country, err := utils.PromptStrInput("Country: ")
	if err != nil || country == "" {
		logger.PrintlnColoredText("❗Invalid input for Country.", logger.ErrorColor)
		return
	}

	// Create address and contact info
	address := entities.Address{
		Street:     street,
		City:       city,
		State:      state,
		PostalCode: postalCode,
		Country:    country,
	}

	contactInfo := entities.ContactInfo{
		Email: email,
		Phone: phone,
	}

	// Create customer
	newCustomer, err := entities.NewCustomer(customerType, firstName, lastName, companyName, contactInfo, address)
	if err != nil {
		logger.PrintlnColoredText(fmt.Sprintf("❗Cannot create the customer object: %v", err), logger.ErrorColor)
		return
	}

	_, err = h.customerService.Create(newCustomer, ctx)
	if err != nil {
		logger.PrintlnColoredText(fmt.Sprintf("❗Failed to save customer: %v", err), logger.ErrorColor)
		return
	}

	logger.PrintlnColoredText("✅ Customer created successfully!", logger.SuccessColor)
}

func (h *AdminCustomerMgmtHandler) handleUpdateCustomer(ctx context.Context) {

}

func (h *AdminCustomerMgmtHandler) handleDeleteCustomer(ctx context.Context) {
	customerId, err := utils.PromptStrInput("Enter the customer ID: ")
	if err != nil {
		logger.PrintlnColoredText("❗ Invalid input. Please enter valid product ID ❗", logger.ErrorColor)
		return
	}

	err = h.customerService.DeleteById(customerId, ctx)
	if err != nil {
		fmt.Printf("\n%v", err)
	}
}

func (h *AdminCustomerMgmtHandler) handleGetCustomerDetails(ctx context.Context) {
	customerId, err := utils.PromptStrInput("Enter the Customer ID: ")
	if err != nil {
		logger.PrintlnColoredText("❗ Invalid input. Please enter valid product ID ❗", logger.ErrorColor)
		return
	}

	customerDetails, err := h.customerService.GetById(ctx, customerId)
	if err != nil {
		logger.PrintlnColoredText(fmt.Sprintf("❗Error. Cannot get product details [ID=%s]. Error: %v", customerId, err), logger.ErrorColor)
		return
	}

	views.DisplayCustomerDetails(customerDetails)

	fmt.Println("\nPress Enter to continue...")
	_, err = fmt.Scanln()
	if err != nil {
		return
	}
}

func (h *AdminCustomerMgmtHandler) handleSeedCustomers(ctx context.Context) {
	err := h.customerService.Seed(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

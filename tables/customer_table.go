package tables

import (
	"fmt"
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ListCustomers(customers []entities.Customer) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, "WSC - Customers 📦")
	table.SetHeader([]string{"ID", "Name", "Type"})

	for _, customer := range customers {
		var idString string
		if customer.ID.IsZero() {
			idString = "N/A"
		} else {
			idString = customer.ID.Hex()
		}

		name := ""
		if customer.CompanyName == "" && customer.CustomerType == entities.Company {
			name = customer.FirstName + ", " + customer.LastName
		} else {
			name = customer.CompanyName
		}

		var customerType string
		if customer.CustomerType == entities.Company {
			customerType = "Company"
		} else {
			customerType = "Individual"
		}

		table.Append([]string{
			idString,
			name,
			customerType,
		})
	}

	table.SetFooter([]string{"", "", "Count", fmt.Sprintf("%d", len(customers))})
	table.Render()
}

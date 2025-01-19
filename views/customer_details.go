package views

import (
	"github.com/Sinet2000/go-eshop-console/internal/entities"
	"github.com/olekukonko/tablewriter"
	"os"
)

func DisplayCustomerDetails(customer *entities.Customer) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetBorder(true)

	// Add general customer details
	table.Append([]string{"ID", customer.ID.Hex()})
	table.Append([]string{"First Name", customer.FirstName})
	table.Append([]string{"Last Name", customer.LastName})
	if customer.CustomerType == entities.Company {
		table.Append([]string{"Company Name", customer.CompanyName})
	}
	table.Append([]string{"Customer Type", map[bool]string{true: "Company", false: "Individual"}[customer.CustomerType == entities.Company]})

	// Add contact information
	table.Append([]string{"Email", customer.ContactInfo.Email})
	table.Append([]string{"Phone", customer.ContactInfo.Phone})

	// Add address details
	table.Append([]string{"Street", customer.Address.Street})
	table.Append([]string{"City", customer.Address.City})
	table.Append([]string{"State", customer.Address.State})
	table.Append([]string{"Postal Code", customer.Address.PostalCode})
	table.Append([]string{"Country", customer.Address.Country})

	// Render the table
	table.Render()
}

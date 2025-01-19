package views

import (
	"fmt"
	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

func DisplayAdminCustomerMgmtMenu() {
	fmt.Println("WSC - Admin Customer Management Menu 🛠️")
	currentTime := time.Now().Format("2006-01-02 15:04")
	fmt.Printf("Hello %s - %s\n", config.GetEnv("ADMIN_NAME"), currentTime)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", ""})

	table.Append([]string{"[1]", "List All Customers"})
	table.Append([]string{"[2]", "List Customers (Paginated)"})
	table.Append([]string{"[3]", "Create Customer"})
	table.Append([]string{"[4]", "Update Customer"})
	table.Append([]string{"[5]", "Delete Customer"})
	table.Append([]string{"[6]", "Get Customer details"})
	table.Append([]string{"[9]", "Seed Customers"})
	table.Append([]string{"[0]", "Quit 🚪"})

	table.Render()
}

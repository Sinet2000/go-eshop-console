package views

import (
	"fmt"
	"os"
	"time"

	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/olekukonko/tablewriter"
)

func DisplayAdminProductMngmtMenu() {
	fmt.Println("WSC - Admin Menu 🛠️")
	currentTime := time.Now().Format("2006-01-02 15:04")
	fmt.Printf("Hello %s - %s\n", config.GetEnv("ADMIN_NAME"), currentTime)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", ""})

	table.Append([]string{"[1]", "List All Products"})
	table.Append([]string{"[2]", "List Products (Paginated)"})
	table.Append([]string{"[3]", "Create Product"})
	table.Append([]string{"[4]", "Update Product"})
	table.Append([]string{"[5]", "Delete Product"})
	table.Append([]string{"[0]", "Quit 🚪"})

	table.Render()
}

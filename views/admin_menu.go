package views

import (
	"fmt"
	"os"
	"time"

	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/olekukonko/tablewriter"
)

func DispalyAdminMenu() {
	fmt.Println("WSC - Admin Menu 🛠️")
	currentTime := time.Now().Format("2006-01-02 15:04")
	fmt.Printf("Hello %s - %s\n", config.GetEnv("ADMIN_NAME"), currentTime)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", ""})

	table.Append([]string{"[1]", "Manage Products 📦"})
	table.Append([]string{"[2]", "Manage Orders 🛒"})
	table.Append([]string{"[3]", "Manage Customers 👥"})
	table.Append([]string{"[4]", "Analytics 📊"})
	table.Append([]string{"[5]", "System Settings ⚙️"})
	table.Append([]string{"[0]", "Quit 🚪"})
	table.Render()
}

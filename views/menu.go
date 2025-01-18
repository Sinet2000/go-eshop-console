package views

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func DisplayMainMenu() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Option", "Action"})
	// table.SetBorder(true)
	// table.SetHeaderLine(true)
	// table.SetCenterSeparator("|")

	// Add menu items with emojis in the action column
	table.Append([]string{"[1]", "List Products 📜"})
	table.Append([]string{"[2]", "Product Details 📝"})
	table.Append([]string{"[3]", "Edit Product 🔄"})
	table.Append([]string{"[4]", "Delete Product 🗑️"})
	table.Append([]string{"[5]", "Create Product 🆕"})
	table.Append([]string{"[6]", "Seed Products 🆕"})
	table.Append([]string{"[0]", "Quit 🛑"})
	table.Render() // Print the formatted table
}

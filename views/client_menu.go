package views

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func DisplayClientMenu() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Option", "Action"})

	table.Append([]string{"[1]", "Go Shopping🛒"})
	table.Append([]string{"[0]", "Quit 🛑"})
	table.Render()
}
